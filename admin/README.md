# admin

- antd
- nextjs
- serverless deploy
- don't need to use serverless
- google authentication

# serverless-nextjs-plugin 걷어낼 예정

lambda@edge + cloudfront 이 구조임, lambda@edge인건 좋은데, 이 .. nextjs 플러그인의 라우팅이 블랙박스라서 짜증남

- 문제1 trigger하기가 어려움, 코드가 어딴식으로 짜져있는지 파악이 어려움. 그렇기에 CloudFront의 이벤트 데이터 타입인 request-origin 메세지를 만들기가 애매함. 레퍼런스 찾기도 힘들고 연구하자니 짜증남.
- 문제2 문제1로인해 console.log를 직어도 해당 api에 도달을 못함. 그리고 실제 e2e로 테스트해볼라해도 이게 캐싱이 되서그런건지, 확실하지 않기에 디버깅환경이 진짜 개똥같음. 개발은 편하겠지...
- 문제3 문제1과 문제2가 있어도 로깅을 어떻게든 달아넣으려고하면 할 수 있다고 함.AWS CloudFront에 실시간 로그 기능 달아넣으려고 kinesis를 사용하는데, kinesis + kinesis datastream + cloudwatch 개발공수보다 월에 6만원 이상나가서 비쌈. ㅅㅂ 그럴거면 ECS써서 편하게 log찍고말지 kinesis 샤드 놓고만있어도 하루에 1400원에서 2000원씩 나감.
