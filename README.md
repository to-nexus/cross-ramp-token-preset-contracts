# Cross-Ramp Token Preset Contracts
This repository provides preset smart contracts for ERC20, ERC721, and ERC1155 tokens, designed for flexible and secure minting operations. The contracts are structured to allow easy extension and integration with custom business logic.

## Key Features
Token Types: Supports ERC20, ERC721, and ERC1155 standards.
Forge System: Each token contract maintains a list of authorized "Forge" contracts. Only addresses in this list can execute minting functions, ensuring controlled token issuance.
TokenBase Management: The TokenBase abstract contract manages the list of forges and provides utility functions for forge management. All token types inherit from their respective Base contracts, which in turn inherit from TokenBase.
Preset Implementations: For each token standard, there are simple preset contracts (e.g., ERC20Mintable, ERC721Simple, ERC1155Simple) that can be deployed directly or extended for custom use cases.
ERC20 Extensions: The ERC20 implementation includes additional extension contracts (such as capped supply, initial supply, mint limits, and period-based minting) to support a wide range of tokenomics.
## Usage
To create a new token, inherit from the appropriate Base contract (ERC20Base, ERC721Base, or ERC1155Base) and implement your custom logic as needed.
Manage the list of authorized forges using the provided functions in TokenBase.
Use the preset contracts for quick deployment or as templates for further customization.

<br><br>
---

# Cross-Ramp 토큰 프리셋 컨트랙트
이 저장소는 ERC20, ERC721, ERC1155 표준을 지원하는 토큰 프리셋 스마트 컨트랙트를 제공합니다. 유연하고 안전한 민팅 기능을 위해 설계되었으며, 비즈니스 로직에 맞게 쉽게 확장할 수 있습니다.

## 주요 특징
토큰 타입: ERC20, ERC721, ERC1155 표준 지원
Forge 시스템: 각 토큰 컨트랙트는 민팅 권한이 있는 "Forge" 컨트랙트 목록을 관리합니다. 이 목록에 포함된 주소만 민팅 함수를 실행할 수 있어, 토큰 발행을 안전하게 제어할 수 있습니다.
TokenBase 관리: TokenBase 추상 컨트랙트가 forge 목록을 관리하며, forge 관련 유틸리티 함수를 제공합니다. 모든 토큰 타입은 각자의 Base 컨트랙트를 상속받아 구현됩니다.
프리셋 구현: 각 토큰 표준별로 간단한 프리셋 컨트랙트(ERC20Mintable, ERC721Simple, ERC1155Simple 등)가 제공되어 바로 배포하거나 커스터마이징의 템플릿으로 활용할 수 있습니다.
ERC20 확장 기능: ERC20 구현에는 공급량 제한, 초기 공급, 민팅 제한, 기간별 민팅 등 다양한 확장 기능이 포함되어 있어 다양한 토크노믹스에 대응할 수 있습니다.
## 사용법
새로운 토큰을 만들 때는 해당 Base 컨트랙트(ERC20Base, ERC721Base, ERC1155Base)를 상속받아 필요한 커스텀 로직을 구현하면 됩니다.
TokenBase에서 제공하는 함수를 통해 forge 목록을 관리할 수 있습니다.
프리셋 컨트랙트를 활용해 빠르게 배포하거나, 추가 개발의 템플릿으로 사용할 수 있습니다.