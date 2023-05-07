import rlp 

def main():
    print(rlp.encode(0x1))
    print(rlp.encode(0x7f))
    print(rlp.encode("ab"))
    print(rlp.encode("ab").hex())
                     

if __name__ == "__main__":
    main()