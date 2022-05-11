/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../dex/params";

export const protobufPackage = "soupyfinance.noodle.dex";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryBookRequest {
  market: string;
  side: boolean;
}

export interface QueryBookResponse {
  pure: string;
  amm: string;
}

export interface QueryBooksRequest {
  market: string;
}

export interface QueryBooksResponse {
  bids: string;
  asks: string;
  amm: string;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryBookRequest: object = { market: "", side: false };

export const QueryBookRequest = {
  encode(message: QueryBookRequest, writer: Writer = Writer.create()): Writer {
    if (message.market !== "") {
      writer.uint32(10).string(message.market);
    }
    if (message.side === true) {
      writer.uint32(16).bool(message.side);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBookRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBookRequest } as QueryBookRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.market = reader.string();
          break;
        case 2:
          message.side = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBookRequest {
    const message = { ...baseQueryBookRequest } as QueryBookRequest;
    if (object.market !== undefined && object.market !== null) {
      message.market = String(object.market);
    } else {
      message.market = "";
    }
    if (object.side !== undefined && object.side !== null) {
      message.side = Boolean(object.side);
    } else {
      message.side = false;
    }
    return message;
  },

  toJSON(message: QueryBookRequest): unknown {
    const obj: any = {};
    message.market !== undefined && (obj.market = message.market);
    message.side !== undefined && (obj.side = message.side);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBookRequest>): QueryBookRequest {
    const message = { ...baseQueryBookRequest } as QueryBookRequest;
    if (object.market !== undefined && object.market !== null) {
      message.market = object.market;
    } else {
      message.market = "";
    }
    if (object.side !== undefined && object.side !== null) {
      message.side = object.side;
    } else {
      message.side = false;
    }
    return message;
  },
};

const baseQueryBookResponse: object = { pure: "", amm: "" };

export const QueryBookResponse = {
  encode(message: QueryBookResponse, writer: Writer = Writer.create()): Writer {
    if (message.pure !== "") {
      writer.uint32(10).string(message.pure);
    }
    if (message.amm !== "") {
      writer.uint32(18).string(message.amm);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBookResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBookResponse } as QueryBookResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pure = reader.string();
          break;
        case 2:
          message.amm = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBookResponse {
    const message = { ...baseQueryBookResponse } as QueryBookResponse;
    if (object.pure !== undefined && object.pure !== null) {
      message.pure = String(object.pure);
    } else {
      message.pure = "";
    }
    if (object.amm !== undefined && object.amm !== null) {
      message.amm = String(object.amm);
    } else {
      message.amm = "";
    }
    return message;
  },

  toJSON(message: QueryBookResponse): unknown {
    const obj: any = {};
    message.pure !== undefined && (obj.pure = message.pure);
    message.amm !== undefined && (obj.amm = message.amm);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBookResponse>): QueryBookResponse {
    const message = { ...baseQueryBookResponse } as QueryBookResponse;
    if (object.pure !== undefined && object.pure !== null) {
      message.pure = object.pure;
    } else {
      message.pure = "";
    }
    if (object.amm !== undefined && object.amm !== null) {
      message.amm = object.amm;
    } else {
      message.amm = "";
    }
    return message;
  },
};

const baseQueryBooksRequest: object = { market: "" };

export const QueryBooksRequest = {
  encode(message: QueryBooksRequest, writer: Writer = Writer.create()): Writer {
    if (message.market !== "") {
      writer.uint32(10).string(message.market);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBooksRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBooksRequest } as QueryBooksRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.market = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBooksRequest {
    const message = { ...baseQueryBooksRequest } as QueryBooksRequest;
    if (object.market !== undefined && object.market !== null) {
      message.market = String(object.market);
    } else {
      message.market = "";
    }
    return message;
  },

  toJSON(message: QueryBooksRequest): unknown {
    const obj: any = {};
    message.market !== undefined && (obj.market = message.market);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBooksRequest>): QueryBooksRequest {
    const message = { ...baseQueryBooksRequest } as QueryBooksRequest;
    if (object.market !== undefined && object.market !== null) {
      message.market = object.market;
    } else {
      message.market = "";
    }
    return message;
  },
};

const baseQueryBooksResponse: object = { bids: "", asks: "", amm: "" };

export const QueryBooksResponse = {
  encode(
    message: QueryBooksResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.bids !== "") {
      writer.uint32(10).string(message.bids);
    }
    if (message.asks !== "") {
      writer.uint32(18).string(message.asks);
    }
    if (message.amm !== "") {
      writer.uint32(26).string(message.amm);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBooksResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBooksResponse } as QueryBooksResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.bids = reader.string();
          break;
        case 2:
          message.asks = reader.string();
          break;
        case 3:
          message.amm = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBooksResponse {
    const message = { ...baseQueryBooksResponse } as QueryBooksResponse;
    if (object.bids !== undefined && object.bids !== null) {
      message.bids = String(object.bids);
    } else {
      message.bids = "";
    }
    if (object.asks !== undefined && object.asks !== null) {
      message.asks = String(object.asks);
    } else {
      message.asks = "";
    }
    if (object.amm !== undefined && object.amm !== null) {
      message.amm = String(object.amm);
    } else {
      message.amm = "";
    }
    return message;
  },

  toJSON(message: QueryBooksResponse): unknown {
    const obj: any = {};
    message.bids !== undefined && (obj.bids = message.bids);
    message.asks !== undefined && (obj.asks = message.asks);
    message.amm !== undefined && (obj.amm = message.amm);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBooksResponse>): QueryBooksResponse {
    const message = { ...baseQueryBooksResponse } as QueryBooksResponse;
    if (object.bids !== undefined && object.bids !== null) {
      message.bids = object.bids;
    } else {
      message.bids = "";
    }
    if (object.asks !== undefined && object.asks !== null) {
      message.asks = object.asks;
    } else {
      message.asks = "";
    }
    if (object.amm !== undefined && object.amm !== null) {
      message.amm = object.amm;
    } else {
      message.amm = "";
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Book items. */
  Book(request: QueryBookRequest): Promise<QueryBookResponse>;
  /** Queries a list of Books items. */
  Books(request: QueryBooksRequest): Promise<QueryBooksResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "soupyfinance.noodle.dex.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Book(request: QueryBookRequest): Promise<QueryBookResponse> {
    const data = QueryBookRequest.encode(request).finish();
    const promise = this.rpc.request(
      "soupyfinance.noodle.dex.Query",
      "Book",
      data
    );
    return promise.then((data) => QueryBookResponse.decode(new Reader(data)));
  }

  Books(request: QueryBooksRequest): Promise<QueryBooksResponse> {
    const data = QueryBooksRequest.encode(request).finish();
    const promise = this.rpc.request(
      "soupyfinance.noodle.dex.Query",
      "Books",
      data
    );
    return promise.then((data) => QueryBooksResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
