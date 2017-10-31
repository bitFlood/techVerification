def hello() :
    str_v = "SEND 26738498932892398 TO 278134903287127384"

    #HARD CODING - right message must have 4 splits.
    strs = str_v.split(' ')

    if not (len(strs) == 4 and strs[0] == 'SEND' and strs[2] == 'TO') :
        return None

    return (strs[1], strs[3])

if __name__ == "__main__" :
    print(hello())