import rlp 

def main():
    print(rlp.encode(0x1))
    print(rlp.encode(0x7f))
    print(rlp.encode("ab")) # 리턴값 b'\x82ab', 0x80 + 문자열길이 2 = 0x82에다가 a와 b 각각을 붙임
    print(rlp.encode("ab").hex()) #리턴값 '826162', 8+2+a의 아스키값(61), b의 아스키값(62)
    print(rlp.encode("hello")) # 리턴값 b'\x85hello' 0x80+문자열길이(5), 거기에 h+e+l+l+o 붙임


if __name__ == "__main__":
    main()