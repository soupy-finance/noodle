/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "soupyfinance.noodle.oracle";

/** Params defines the parameters for the module. */
export interface Params {
  assets: string[];
}

const baseParams: object = { assets: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    for (const v of message.assets) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    message.assets = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.assets.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    message.assets = [];
    if (object.assets !== undefined && object.assets !== null) {
      for (const e of object.assets) {
        message.assets.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    if (message.assets) {
      obj.assets = message.assets.map((e) => e);
    } else {
      obj.assets = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    message.assets = [];
    if (object.assets !== undefined && object.assets !== null) {
      for (const e of object.assets) {
        message.assets.push(e);
      }
    }
    return message;
  },
};

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
