const mod = Process.getModuleByName("WeChat");
console.log("[+] Base Address:", mod.base);

function checkValid(p) {
    if (p.isNull()) {
        return false;
    }

    if (!p.and(0x7).isNull()) {
        return false;
    }
    if (p.compare(ptr("0x600000000000")) >= 0 && p.compare(ptr("0x700000000000")) < 0) {
        return true;
    }
    return false;
}

function handlePr(addr, keyword) {
    const realAddr = ptr("0x" + addr).sub("0x100000000").add(mod.base);
    console.log("[+] real Address:", realAddr)

    Interceptor.attach(realAddr, {
        onEnter(args) {
            for (let i = 0; i < 10; i++) {
                try {
                    if (checkValid(args[i])) {
                        const buf = args[i].readByteArray(128)
                        if (!buf) {
                            continue;
                        }
                        let s = "";
                        const u8 = new Uint8Array(buf);
                        for (let b of u8) {
                            if (b >= 0x20 && b <= 0x7E) {
                                s += String.fromCharCode(b);
                            } else {
                                s += ".";
                            }
                        }

                        if (s.includes(keyword) || keyword === "") {
                            console.log(`\n[+] ${addr} || arg${i} ${args[i]}`);
                            console.log(hexdump(args[i], {length: 64}));
                            console.log(addr + "|| " + s + "\n" +
                                Thread.backtrace(this.context, Backtracer.ACCURATE)
                                    .map(DebugSymbol.fromAddress).join('\n'));
                            return;
                        }
                    }
                } catch (e) {
                    console.log("Enter Error:", e);
                }
            }


        },

        onLeave(retval) {
            console.log(`===== ${addr} LEAVE =====`);
            console.log(` ${addr}  Return value ${retval}`);
            try {
                if (checkValid(retval)) {
                    console.log(addr + "||" + hexdump(retval, {
                        offset: 0,
                        length: 64
                    }));
                }

            } catch (_) {
            }
        }
    });

}


const prs = ["104565AD0", "104566E24", "1045CF820", "104590888", "1045BFED0", "104394290", "1043877CC", "104387764", "1043382C0"]
const k = "";
for (let pr of prs) {
    handlePr(pr, k);
}

