# 注意

libcomm__xlogger.cc.sqlite 
libcomm__xloggerbase.c.sqlite 
用这个sqlite文件


通过关键字 DoTypeSafeFormat 定位到sub_1046385AC打日志函数

下面准备解字符串

unk_1083E89B8 有下面的常量
1083D1F88 aEarlierNotific 常量，怎么使用

108260910 -> 108260918 -> unk_1083E89B8 -> 1083E89C8 -> 1083D1F88

AddMessageSendContext 关键字 1024C10E8 函数

需要先发一条消息
sub_102480484 -> checkPrepareShowMsg_102482A74 -> sub_1024C4CB4
sub_102480484->
StartSendMessageSerial sub_1024C4CB4 发消息的函数
CoSendMessageWithUploadInfo sub_1023E8108
CoAddSendMessageToDb  sub_1023C09D0
CoPrepareShowSendMessage  sub_1023BC4E0
sendfinish sub_1024C7FB4
具体发的函数：sub_102481CA0


v81[0]  // 自引用
v81[1]  // 函数sub_1018D7A90
v81[2]  // 函数 sub_1023C25F0
v81[3]  // 数字1 可能是消息类型    
v81[4]  // 连续指针，最后值为空指针 
v81[5]  // 连续指针，最后值为空指针
v81[6]  // 对齐字段 0x10000000
v81[7]  // 连续指针，最后值为空指针
v81[8]  // Begin StartSend Message SyncStag
v81[9]  //数字，可能是消息id 0x304400018B237EB4      
v81[10] // 空指针
v81[11] // 空指针
v81[12]// 空指针
v81[13] // 空指针
v81[14] // 指针，值为空
v81[15] // 指针，值为空
v81[16] // 指针，值为空
v81[17] //数字，可能是消息id
v81[18] // 指针，值为空
v81[19] // 指针，值为空
v81[20] // 指针，值为空
v81[21] // 数字
v81[22] // 指针，值为空
v81[23] // 指针，值为空
v81[24] // Begin StartSend Message SyncStag
v81[25] // sub_1021D5E48 函数地址
v81[26] // sub_10250D544 函数地址
v81[27] // sub_10247C008  函数
v81[28] // 连续指针，最后值为空指针
v81[29] = 0x000000010250DB70      // 175ED6928