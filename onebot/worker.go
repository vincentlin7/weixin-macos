package main

import (
	"context"
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

func SendWorker() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("💥 SendWorker 异常: %v\n", err)
			go SendWorker()
		}
	}()
	
	for {
		select {
		case <-finishChan:
			fmt.Printf("收到完成信号 \n")
		case m, ok := <-msgChan:
			if !ok {
				return
			}
			SendWechatMsg(m)
		}
	}
}

func SendWechatMsg(m *SendMsg) {
	time.Sleep(1 * time.Second)
	currTaskId := atomic.AddInt64(&taskId, 1)
	log.Printf("📩 收到任务: %d\n", currTaskId)
	
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	
	targetId := m.UserId
	if m.GroupID != "" && targetId == "" {
		targetId = m.GroupID
	}
	
	switch m.Type {
	case "text":
		result := fridaScript.ExportsCall("triggerSendTextMessage", currTaskId, targetId, m.Content, m.AtUser)
		log.Printf("📩 发送文本任务执行结果：%s, 参数：currTaskId: %d, targetId: %s, content: %s, atUser: %s\n",
			result, currTaskId, targetId, m.Content, m.AtUser)
	case "image":
		targetPath, md5Str, err := SaveBase64Image(m.Content)
		if err != nil {
			log.Printf("保存图片失败: %v\n", err)
			return
		}
		
		result := fridaScript.ExportsCall("triggerUploadImg", targetId, md5Str, targetPath)
		log.Printf("📩 上传图片任务执行结果%s, 参数：targetId: %s, md5Str: %s, targetPath: %s\n", result, targetId, md5Str, targetPath)
	case "send_image":
		result := fridaScript.ExportsCall("triggerSendImgMessage", currTaskId, myWechatId, targetId)
		log.Printf("📩 发送图片任务执行结果%s, 参数：currTaskId: %d, myWechatId: %s, targetId: %s\n", result, currTaskId, myWechatId, targetId)
	}
	
	select {
	case <-ctx.Done():
		log.Printf("任务 %d 执行超时！\n", currTaskId)
	case <-finishChan:
		log.Printf("收到完成信号，任务 %d 完成\n", currTaskId)
	}
}
