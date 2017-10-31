'''
2017 / 10 / 31 Juung Bae
머클트리의 hash를 만드는 알고리즘 테스트 코드
@logic @muckle @hash
'''
import hashlib
import time

class Node:
    childs = []
    hash = ""

    def __repr__(self) :
        return self.hash

    def __init__(self, *args):
        if len(args) <= 0 :
            hash = hashlib.sha256(str(round(time.time()*1000)).encode('utf-8')).hexdigest()
        
        else :
            string = ""
            for i in args :
                self.childs.append(i)
                string = string + repr(i)
            
            self.hash = hashlib.sha256(string.encode('utf-8')).hexdigest()

input_data = []
curr_datas = [input_data]
p = 0

#Initial
for i in range(0, 7) :
    input_data.append(Node)

print(input_data)

while True :
    new_data = []
    for i in range(0, round(len(curr_datas[p])/2)):
        if i+1 >= len(curr_datas) :
            new_data.append(Node(curr_datas[p][i]))
        else :
            new_data.append(Node(curr_datas[p][i], curr_datas[p][i+1]))

    curr_datas.append(new_data)
    print(new_data)
    if len(new_data) <= 1 :
        break
    p = p + 1