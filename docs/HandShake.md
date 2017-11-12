# BasicChain 체인 Handshake 과정

> 2017.11.12
>
> BasicChain 노드 서버 작성 전 체인 Handshake에 대한 간략한 명세입니다.



## 새로운 블록 생성

* ip테이블에 있는 다른 노드들에게 새로운 블럭을 체인에 연결하기 전에 업데이트 할 사항을 받아온다.
  * **(있을때)** 블록을 파기하고, 체인을 업데이트한다.
  * **(없을때)** 새로 생성된 블록을 체인에 연결하고 다른 노드들에게 전파한다.



## 블록 업데이트

* 다른 노드에게서 블록 전파 요청이 들어온다.
  * 새로운 블록 데이터를 받아오고, 유효한 블록인지 증명과정을 거친다.
  * 유효하면, 체인에 연결하고 승인 신호 전달.



## 블록 Proof

* 블록체인을 업데이트할때, 받아오는 새로운 블록이 유효하고, 신뢰성이 있는 블록인지 확인하는 작업.
  * dPoS (delegated Proof of Stake)
  * PoS (Proof of Stake)
  * **PoS (Proof of Space)**
  * PoW (Proof of Work)
  * 등이 있는데, Basic Chain에서는 작업 증명 안할거 ㅋ