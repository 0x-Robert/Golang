import rlp 

#https://docs.google.com/presentation/d/1nSoRv4hCmona_N1VENZdZ_POVV-LZpxOSNBntareZuA/edit#slide=id.g3597ea7b61_0_91
def main():
    print(rlp.encode(0x1))
    print(rlp.encode(0x7f))
    print(rlp.encode("ab")) # 리턴값 b'\x82ab', 0x80 + 문자열길이 2 = 0x82에다가 a와 b 각각을 붙임
    print(rlp.encode("ab").hex()) #리턴값 '826162', 8+2+a의 아스키값(61), b의 아스키값(62)
    print(rlp.encode("hello")) # 리턴값 b'\x85hello' 0x80+문자열길이(5), 거기에 h+e+l+l+o 붙임
    
    # 리턴값 <b'\xb9\x04\x00aaaa….a가 2014개 >
    #설명하자면, 1024를 바이트로 표현하면, <0x04> <0x00> 2바이트로 표현된다.
    #따라서 <0xb7 + 0x02> 뒤에 0x04 0x00을 붙이면 → 0xb9 0x04 0x00 → 위의 <> 앞자리와 같음
    #빨간색은 0xb7에 2를 더한것이고(길이가 2바이트로 표현되었기 때문)
    #파란색은 1024를 바이트로 표현한 것이다.(1024 → 0x04 0x00)
    print(rlp.encode("a"*1024)) 
    print(rlp.encode("a"*91))

    # 리턴값은 b'\xc3\x82ab', “ab”를 RLP인코딩하면, 0x82 a b 가 되고(rlp definition 1 참조),
    # 이는 3자리 이므로, 0xc0+3 = 0xc3, 따라서 최종 값은 <0xc3 0x82 a b> → b'\xc3\x82ab'
    print(rlp.encode(["ab"])) 
    
    #ab를 인코딩하면 0x82 a b, “dc”를 인코딩하면 0x82 d c, 둘의 길이를 합치면 6 따라서
    #0xc0+6 뒤에 0x82 a b, 0x82 d c → b'\xc6\x82ab\x82ab'
    print(rlp.encode(["ab","dc"]))

    print("test")
    print(rlp.encode("adb")) #b'\x83adb'
    print(rlp.encode(11))    #b'\x0b'
    print(rlp.encode(["adb", 0x11])) #b'\xc5\x83adb\x11'
    print(rlp.encode(["a"*50, "a"*50])) #b'\xf8f\xb2aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\xb2aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa'
    


if __name__ == "__main__":
    main()