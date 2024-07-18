## Go 의 기본적인 문법, 특징 공부

### something

- Go 에서 function 을 export 하는 방법
  - 함수의 첫 글자를 대문자로 하면 export 된다.

### Basic

- Go 의 기본적인 문법
  - 변수 선언
    - 배열 선언
  - 상수 선언
  - 함수 선언
    - naked return
      - 함수의 리턴값을 여러개로 반환할 수 있다.
    - defer
      - 함수가 종료되기 직전에 실행되는 함수
      - stack 처럼 먼저 선언된 defer 가 나중에 실행된다. 
  - 조건문
  - 반복문
  - 전개 문법은 A... 으로 사용한다.
    - javascript 는 ...A 로 사용

### Special

- pointer
- append
- map
- struct