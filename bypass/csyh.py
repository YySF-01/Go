def xor(shellcode, key):
    new_shellcode = ""
    key_len = len(key)
    # 对shellcode的每一位进行xor亦或处理
    for i in range(0, len(shellcode)):
        s = ord(shellcode[i])
        p = ord((key[i % key_len]))
        s = s ^ p  # 与p异或，p就是key中的字符之一
        s = chr(s) 
        new_shellcode += s
    return new_shellcode

def random_decode(shellcode):
    j = 0
    new_shellcode = ""
    for i in range(0,len(shellcode)):
        if i % 2 == 0:
            new_shellcode[i] = shellcode[j]
            j += 1

    return new_shellcode

def add_random_code(shellcode, key):
    new_shellcode = ""
    key_len = len(key)
    # 每个字节后面添加随机一个字节，随机字符来源于key
    for i in range(0, len(shellcode)):
        #print(ord(shellcode[i]))
        new_shellcode += shellcode[i]
        # print("&"+hex(ord(new_shellcode[i])))
        new_shellcode += key[i % key_len]

        #print(i % key_len)
    return new_shellcode

# 将shellcode打印输出
def str_to_hex(shellcode):
    raw = ""
    for i in range(0, len(shellcode)):
        s = hex(ord(shellcode[i])).replace("0x",',0x')
        raw = raw + s
    return raw

if __name__ == '__main__':
    shellcode = ""
    # 这是异或和增加随机字符使用的key
    key = "iqe"
    print(shellcode[0])
    print(len(shellcode))
    # 首先对shellcode进行异或处理
    shellcode = xor(shellcode, key)
    print(len(shellcode))

    # 然后在shellcode中增加随机字符
    shellcode = add_random_code(shellcode, key)

    # 将shellcode打印出来
    print(str_to_hex(shellcode))