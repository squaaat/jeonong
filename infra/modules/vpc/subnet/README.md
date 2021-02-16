#### 최소한 보고와야하는 문서

- https://docs.aws.amazon.com/ko_kr/vpc/latest/userguide/VPC_Scenario2.html
- https://www.youtube.com/watch?v=hiKPPy584Mg

# Subnet이용가이드

## `public`
- Public subnet
- attached internet gateway to route table
- have NAT Gateways
- `default Route table row` = Internet Gateway

해당 노드는 public IP를 항상 할당받아서 이용하는 목적으로 생성한 서브넷이다. inbound는 가능한 80, 443같은 가장 일반적인 포트로만 요청을 받는 목적.
internet 망에 1차원적으로 직접 부딛히는 노드들만 접근하는 망.

#### Basic Rules
- `EC2를 직접 개설해서 쓰지 않도록 한다. 인프라 주요멤버 아니면 가능한 쓰지 않도록 한다.`
- 각 public node마다 별도의 security_group를 받을 수 있도록 한다.
- `[sg_members, sg_basic]` 라는 SG들을 반드시 등록해서 사용할 수 있도록 한다.

## `private_nat`

- Private subnet with NAT gateway.
- Don't allow network access from externals.
- But, Only allow access from node on private subents to externals
- `default Route table row` = NAT Gateway

public IP 할당받지 않고 쓰면서, 외부 third party로 부터 API 호출은 하고 싶지만, 외부로 부터 접근 하고 싶지 않은 리소스들은 해당 subnet을 이용할 수 있도록 한다.

#### Basic Rules
- 최소한 `[sg_basic]` 라는 SG들을 반드시 등록해서 사용할 수 있도록 한다.
- 일반적인 application이라면 해당 subnet을 이용할 수 있도록 한다.

## `private`
- Private subent.
- VPC Peering and Allow VPC Endpoints from AWS platforms
- Don't admit all of traffics from external networks.
- DB같은 Storage 타입의 인스턴스가 들어오는 곳.
- `default Route table row` = X

가능한 권장하기로는 private망만 쓰길 바람, AWS resource가 접근조차도 못하기 때문에 peering을 반드시 해서 쓰도록 설계된 망
github이나 dockerhub같은 외부 리소스로부터 다운받아오는것은 용납이 안되는 곳.

예를들어 운영중에 구글 API를 써야한다면, Google VPC Peering을 해당 private subnet의 Route table로 Peering을 직접 해줄게 아니라면
그냥 얌전히 `private_nat` 를 쓰도록 하자. 해당 subnet은 Application서버용이라기보다 아예 저장소 자체를 네트워크 망에서부터 독립시켜서 운영할 때에만 필요함

#### Basic Rules
- `EC2를 직접 개설해서 쓰지 않도록 한다. 인프라 주요멤버 아니면 가능한 쓰지 않도록 한다.`
- 가능한 저장소 성격의 서비스들에만 사용할 수 있도록 하자
- 사용 룰이 명확한 녀석이다.
- Dockerize를 외부에서 dependency들을 더 이상 땡겨오지 않아도 되는 녀석들에 사용되어도 좋다.
