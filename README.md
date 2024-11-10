# Anti-Corruption Layer dummy example
This is an example of an ACL using Go.

### Context fictional
Let's say we're modernizing a legacy e-commerce system, where we need to integrate a modern microservice with a legacy monolithic order management system.  
  
```mermaid
graph TD
    subgraph "Modern Microservice Architecture"
        MS[Modern Order Service]
        ACL[Anti-Corruption Layer]
        TR[Translation Layer]
        AD[Adapter Layer]
    end

    subgraph "Legacy System"
        LS[Legacy Order System]
    end

    Client[Client Application] -->|Modern REST API| MS
    MS -->|Modern Domain Model| ACL
    ACL -->|Translation| TR
    TR -->|Adaptation| AD
    AD -->|Legacy SOAP/XML| LS

    style ACL fill:#f9f,stroke:#333,stroke-width:4px
    style MS fill:#bbf,stroke:#333,stroke-width:2px
    style LS fill:#fbb,stroke:#333,stroke-width:2px
```  
  
### Layers description
Each layer serves a specific purpose:

`ModernOrderService`: Provides clean interface for modern applications  
`LegacyOrderAdapter`: Handles communication with legacy system  
`OrderTranslator`: Manages data format conversions  
`LegacyOrderService `: Represents the external legacy system  