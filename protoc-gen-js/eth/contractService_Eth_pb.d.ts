export class Contract {
  constructor ();
  getSolidity(): string;
  setSolidity(a: string): void;
  getParamsList(): string[];
  setParamsList(a: string[]): void;
  getPublishList(): string[];
  setPublishList(a: string[]): void;
  getPrivate(): string;
  setPrivate(a: string): void;
  getGasLimit(): number;
  setGasLimit(a: number): void;
  getValue(): number;
  setValue(a: number): void;
  getName(): string;
  setName(a: string): void;
  getBin(): string;
  setBin(a: string): void;
  getAbi(): string;
  setAbi(a: string): void;
  getAddress(): string;
  setAddress(a: string): void;
  getCreated(): string;
  setCreated(a: string): void;
  getCreationTxHash(): string;
  setCreationTxHash(a: string): void;
  getResultsList(): string[];
  setResultsList(a: string[]): void;
  toObject(): Contract.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => Contract;
}

export namespace Contract {
  export type AsObject = {
    Solidity: string;
    ParamsList: string[];
    PublishList: string[];
    Private: string;
    GasLimit: number;
    Value: number;
    Name: string;
    Bin: string;
    Abi: string;
    Address: string;
    Created: string;
    CreationTxHash: string;
    ResultsList: string[];
  }
}

export class ContractArray {
  constructor ();
  getContractList(): Contract[];
  setContractList(a: Contract[]): void;
  toObject(): ContractArray.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => ContractArray;
}

export namespace ContractArray {
  export type AsObject = {
    ContractList: Contract[];
  }
}

export class GetContractAddressEndpointRequest {
  constructor ();
  getNetwork(): NetworkAllowingAlias;
  setNetwork(a: NetworkAllowingAlias): void;
  getQueryAddress(): string;
  setQueryAddress(a: string): void;
  toObject(): GetContractAddressEndpointRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => GetContractAddressEndpointRequest;
}

export namespace GetContractAddressEndpointRequest {
  export type AsObject = {
    Network: NetworkAllowingAlias;
    QueryAddress: string;
  }
}

export class PostCallContractMethodEndpointRequest {
  constructor ();
  getNetwork(): NetworkAllowingAlias;
  setNetwork(a: NetworkAllowingAlias): void;
  getQueryAddress(): string;
  setQueryAddress(a: string): void;
  getMethod(): string;
  setMethod(a: string): void;
  getSolidity(): string;
  setSolidity(a: string): void;
  getParamsList(): string[];
  setParamsList(a: string[]): void;
  getPublishList(): string[];
  setPublishList(a: string[]): void;
  getPrivate(): string;
  setPrivate(a: string): void;
  getGasLimit(): number;
  setGasLimit(a: number): void;
  getValue(): number;
  setValue(a: number): void;
  getName(): string;
  setName(a: string): void;
  getBin(): string;
  setBin(a: string): void;
  getAbi(): string;
  setAbi(a: string): void;
  getAddress(): string;
  setAddress(a: string): void;
  getCreated(): string;
  setCreated(a: string): void;
  getCreationTxHash(): string;
  setCreationTxHash(a: string): void;
  getResultsList(): string[];
  setResultsList(a: string[]): void;
  toObject(): PostCallContractMethodEndpointRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => PostCallContractMethodEndpointRequest;
}

export namespace PostCallContractMethodEndpointRequest {
  export type AsObject = {
    Network: NetworkAllowingAlias;
    QueryAddress: string;
    Method: string;
    Solidity: string;
    ParamsList: string[];
    PublishList: string[];
    Private: string;
    GasLimit: number;
    Value: number;
    Name: string;
    Bin: string;
    Abi: string;
    Address: string;
    Created: string;
    CreationTxHash: string;
    ResultsList: string[];
  }
}

export class PostCreateContractEndpointRequest {
  constructor ();
  getNetwork(): NetworkAllowingAlias;
  setNetwork(a: NetworkAllowingAlias): void;
  getSolidity(): string;
  setSolidity(a: string): void;
  getParamsList(): string[];
  setParamsList(a: string[]): void;
  getPublishList(): string[];
  setPublishList(a: string[]): void;
  getPrivate(): string;
  setPrivate(a: string): void;
  getGasLimit(): number;
  setGasLimit(a: number): void;
  getValue(): number;
  setValue(a: number): void;
  getName(): string;
  setName(a: string): void;
  getBin(): string;
  setBin(a: string): void;
  getAbi(): string;
  setAbi(a: string): void;
  getAddress(): string;
  setAddress(a: string): void;
  getCreated(): string;
  setCreated(a: string): void;
  getCreationTxHash(): string;
  setCreationTxHash(a: string): void;
  getResultsList(): string[];
  setResultsList(a: string[]): void;
  toObject(): PostCreateContractEndpointRequest.AsObject;
  serializeBinary(): Uint8Array;
  static deserializeBinary: (bytes: {}) => PostCreateContractEndpointRequest;
}

export namespace PostCreateContractEndpointRequest {
  export type AsObject = {
    Network: NetworkAllowingAlias;
    Solidity: string;
    ParamsList: string[];
    PublishList: string[];
    Private: string;
    GasLimit: number;
    Value: number;
    Name: string;
    Bin: string;
    Abi: string;
    Address: string;
    Created: string;
    CreationTxHash: string;
    ResultsList: string[];
  }
}

