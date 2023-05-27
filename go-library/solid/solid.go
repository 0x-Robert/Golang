package main

/*

단일 책임 원칙 SRP Single Responsibility Principle 		: 모든 객체는 책임을 하나만 져야 한다.
개방 폐쇄 원칙 OCP Open Closed Principle 				: 확장에는 열려 있고 변경에는 닫혀 있다.
리스코프 치환 원칙 LSP liskov substitution principle 	: q(x)를 타입 T의 객체 x에 대해 증명할 수 있는 속성이라하자. 그렇다면 S가 T의 하위 타입이라면 q(y)는 타입S의 객체 y에 대해 증명할 수 있어야한다.
인터페이스 분리원칙 ISP Interface Segregation Principle  : 인터페이스 분리 원칙 : 클라이언트는 자신이 이용하지 않는 메서드에 의존하지 않아야 한다.
의존 관계 역전 원칙 DIP Dependency Inversion Principle	: 의존관계 역전 원칙 : 상위계층이 하위계층에 의존하는 전통적인 의존 관계를 반전시킴으로써 상위계층이 하위계층의 구현으로부터 독립되게 할 수 있다.

나쁜 설계는 상호 결합도가 매우 높고 응집도가 낮다.

*/
