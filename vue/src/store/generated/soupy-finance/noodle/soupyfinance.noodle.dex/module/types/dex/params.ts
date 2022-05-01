/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "soupyfinance.noodle.dex";

/** Params defines the parameters for the module. */
export interface Params {
  markets: { [key: string]: string };
}

export interface Params_MarketsEntry {
  key: string;
  value: string;
}

const baseParams: object = {};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    Object.entries(message.markets).forEach(([key, value]) => {
      Params_MarketsEntry.encode(
        { key: key as any, value },
        writer.uint32(10).fork()
      ).ldelim();
    });
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    message.markets = {};
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          const entry1 = Params_MarketsEntry.decode(reader, reader.uint32());
          if (entry1.value !== undefined) {
            message.markets[entry1.key] = entry1.value;
          }
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
    message.markets = {};
    if (object.markets !== undefined && object.markets !== null) {
      Object.entries(object.markets).forEach(([key, value]) => {
        message.markets[key] = String(value);
      });
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    obj.markets = {};
    if (message.markets) {
      Object.entries(message.markets).forEach(([k, v]) => {
        obj.markets[k] = v;
      });
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    message.markets = {};
    if (object.markets !== undefined && object.markets !== null) {
      Object.entries(object.markets).forEach(([key, value]) => {
        if (value !== undefined) {
          message.markets[key] = String(value);
        }
      });
    }
    return message;
  },
};

const baseParams_MarketsEntry: object = { key: "", value: "" };

export const Params_MarketsEntry = {
  encode(
    message: Params_MarketsEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params_MarketsEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams_MarketsEntry } as Params_MarketsEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params_MarketsEntry {
    const message = { ...baseParams_MarketsEntry } as Params_MarketsEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: Params_MarketsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(object: DeepPartial<Params_MarketsEntry>): Params_MarketsEntry {
    const message = { ...baseParams_MarketsEntry } as Params_MarketsEntry;
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
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
