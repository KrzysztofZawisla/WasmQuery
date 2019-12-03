export interface WasmQueryNode extends Node {
    hasClass(): DOMTokenList;
    hasClass(value: string): boolean;
    hasClass(values: string[]): boolean;
    attr(): NamedNodeMap;
    attr(selector: string): string;
    attr(selectors: string[]): string[];
    attr(selector: string, value: string): string;
    attr(selectors: string[], value: string): string[];
    attr(selector: string, values: string[]): string;
    attr(selectors: string[], values: string[]): string[];
    hide(): string;
    /** 
     * Returns width of html element
     */
    width(): string;
    /**
     * Setups and returns width for html element
     */
    width(value: string): string;
    /**
     * Setups and returns width for html element
     */
    width(values: string[]): string[];
    /** 
     * Returns height of html element
     */
    height(): string;
    /**
     * Setups and returns height for html element
     */
    height(value: string): string;
    /**
     * Setups and returns height for html element
     */
    height(values: string[]): string[];
    /**
     * Returns DOMTokenList of html element
     */
    addClass(): DOMTokenList;
    /**
     * Setups class and returns DOMTokenList for html element
     */
    addClass<T>(value: T | string): DOMTokenList;
    /**
     * Setups classes and returns DOMTokenList for html element
     */
    addClass<T>(values: T | string[]): DOMTokenList;
    /**
     * Returns all css attributes for html element
     */
    css(): CSSStyleDeclaration;
    /**
     * Setups and returns css attributes for html element
     */
    css(selector: string | number): string;
    /**
     * Setups and returns css attributes for html element
     */
    css(selectors: string[] | number[]): string[];
    /**
     * Setups and returns css attributes for html element
     */
    css(selector: string | number, value: string): string;
    /**
     * Setups and returns css attributes for html element
     */
    css(selectors: string[] | number[], value: string): string[];
    css(selector: string | number, values: string[]): string;
    /**
     * Setups and returns css attributes for html element
     */
    css(selectors: string[] | number[], values: string[]): string[];
    val<T>(): T | string;
    val<T>(value: T | string): T | string;
    val<T>(value: T | string[]): T | string;
    on(selector: string, func: Function): null;
};

export interface WasmQueryNodeList {
    [Symbol.iterator](): IterableIterator<WasmQueryNode>;
    /**
     * Returns an array of key, value pairs for every entry in the list.
     */
    entries(): IterableIterator<[number, WasmQueryNode]>;
    /**
     * Returns an list of keys in the list.
     */
    keys(): IterableIterator<number>;
    /**
     * Returns an list of values in the list.
     */
    values(): IterableIterator<WasmQueryNode>;
    len(): number;
    attr(): NamedNodeMap[];
    attr(selector: string): string[];
    attr(selectors: string[]): string[][];
    attr(selector: string, value: string): string[];
    attr(selectors: string[], value: string): string[][];
    attr(selector: string, values: string[]): string[];
    attr(selectors: string[], values: string[]): string[][];
    /**
     * Returns all css attributes for elements of node list
     */
    css(): CSSStyleDeclaration[];
    /**
     * Setups and returns css attributes for elements of node list
     */
    css(selector: string | number): string[];
    /**
     * Setups and returns css attributes for elements of node list
     */
    css(selectors: string[] | number[]): string[][];
    /**
     * Setups and returns css attributes for elements of node list
     */
    css(selector: string | number, value: string): string[];
    /**
     * Setups and returns css attributes for elements of node list
     */
    css(selectors: string[] | number[], value: string): string[][];
    css(selector: string | number, values: string[]): string[];
    /**
     * Setups and returns css attributes for elements of node list
     */
    css(selectors: string[] | number[], values: string[]): string[][];
    /**
     * Returns DOMTokenList of elements of node list
     */
    toggleClass(): DOMTokenList;
    /**
     * Setups, Remove class. Returns DOMTokenList for elements of node list
     */
    toggleClass(value: string): DOMTokenList[];
    /**
     * Setups, Remove classes. Returns DOMTokenList for elements of node list
     */
    toggleClass(value: string[]): DOMTokenList[];
    /**
     * Returns value of elements of node list
     */
    val<T>(): T[] | string[];
    /**
     * Setups and returns value for elements of node list
     */
    val<T>(value: T | string): T[] | string[];
    /**
     * Setups and returns value for elements of node list
     */
    val<T>(value: T | string[]): T[] | string[];
    /** 
     * Returns width of elemnets of node list
     */
    width(): string[];
    /**
     * Setups and returns width for elements of node list
     */
    width(value: string): string[];
    /**
     * Setups and returns width for elements of node list
     */
    width(values: string[]): string[];
    /** 
     * Returns height of elemnets of node list
     */
    height(): string[];
    /**
     * Setups and returns height for elements of node list
     */
    height(value: string): string[];
    /**
     * Setups and returns height for elements of node list
     */
    height(values: string[]): string[];
    /**
     * Returns even elements of node list
     */
    even(): WasmQueryNodeList;
    /**
     * Returns odd elements of node list
     */
    odd(): WasmQueryNodeList;
    /**
     * Returns DOMTokenList of elements of node list
     */
    addClass(): DOMTokenList[];
    /**
     * Setups class and returns DOMTokenList for elements of node list
     */
    addClass<T>(value: T | string): DOMTokenList[];
    /**
     * Setups classes and returns DOMTokenList for elements of node list
     */
    addClass<T>(values: T | string[]): DOMTokenList[];
    /**
     * Returns DOMTokenList of elements of node list
     */
    hasClass(): DOMTokenList[];
    hasClass(value: string): boolean[];
    hasClass(values: string[]): boolean[];
    hide(): string[];
    show(): string[];
    showLegacy(): string[]
    showAsBlock(): string[];
    showAsInline(): string[];
    showAsInlineBlock(): string[];
    showAsInlineFlex(): string[];
    showAsInlineTable(): string[];
    showAsListItem(): string[];
    showAsGrid(): string[];
    showAsFlex(): string[];
    showAsContents(): string[];
    showAsTable(): string[];
};

export declare function registerWasmQuery(): string;
export declare function registerWasmQuery(selector: string): string;
export declare function registerWasmQuery(selectors: string[]): string[];
export declare function $(selector: string): WasmQueryNodeList;
export declare function $(selector: Document): Document;
declare function year(): string;
declare function month(): string;
declare function day(): string;
declare function hour(): string;
declare function minute(): string;
declare function second(): string;
declare function nanosecond(): string;
declare function now(): string;
declare function md4(value: string): string;
declare function md5(value: string): string;
declare function sha1(value: string): string;
declare function sha224(value: string): string;
declare function sha512_224(value: string): string;
declare function sha256(value: string): string;
declare function sha512_256(value: string): string;
declare function sha384(value: string): string;
declare function sha512(value: string): string;
declare function adler32(value: string): string;
declare function fnv1_128(value: string): string;
declare function fnv1a_128(value: string): string;
declare function fnv1_32(value: string): string;
declare function fnv1a_32(value: string): string;
declare function fnv1_64(value: string): string;
declare function fnv1a_64(value: string): string;
declare function keccak256(value: string): string;
declare function keccak512(value: string): string;
declare function sha3_224(value: string): string;
declare function sha3_256(value: string): string;
declare function sha3_384(value: string): string;
declare function sha3_512(value: string): string;
declare function abs(value: number): number;
declare function abs(...values: number[]): number[];
declare function acos(value: number): number;
declare function acos(...values: number[]): number[];
declare function ceil(value: number): number;
declare function floor(value: number): number;
declare function round(value: number): number;
declare function randomInt(): number;
declare function randomInt(value: number): number;
declare function randomInt32(): number;
declare function randomInt32(value: number): number;
declare function randomInt64(): number;
declare function randomInt64(value: number): number;
declare function randomFloat32(): number;
declare function randomFloat32(value: number): number;
declare function randomFloat64(): number;
declare function randomFloat64(value: number): number;
declare function gamma(value: number): number;
declare function NaN(): number;
declare function max(value1: number, value2: number, ...otherValues: number[]): number;
declare function min(value1: number, value2: number, ...otherValues: number[]): number;
declare function pow(value1: number, value2: number, ...otherValues: number[]): number;
declare function sqrt(value: number): number;
declare function pi(): number;
declare function piAsString(): string;
declare function e(): number;
declare function eAsString(): string;
declare function phi(): number;
declare function phiAsString(): string;
declare function ln2(): number;
declare function ln2AsString(): string;
declare function ln10(): number;
declare function ln10AsString(): string;
declare function sqrt2(): number;
declare function sqrt2AsString(): string;
declare function log2E(): number;
declare function log10E(): number;
declare function maxInt8(): number;
declare function minInt8(): number;
declare function maxInt16(): number;
declare function minInt16(): number;
declare function maxInt32(): number;
declare function minInt32(): number;
declare function sqrtPi(): number;
declare function sqrtPiAsString(): string;
declare function log(value: number): number;
declare function log10(value: number): number;
declare function inf(sign: number): number;
declare function isNaN(value: number): boolean;
declare function cosh(value: number): number;
declare function j0(value: number): number;
declare function j1(value: number): number;
declare function erf(value: number): number;
declare function erfinv(value: number): number;
declare function erfc(value: number): number;
declare function erfcinv(value: number): number;
declare function exp(value: number): number;
declare function exp2(value: number): number;
declare function expm1(value: number): number;
declare function hypot(value1: number, value2: number): number;
declare function ilogb(value: number): number;
declare function y0(value: number): number;
declare function y1(value: number): number;
declare function tanh(value: number): number;
declare function trunc(value: number): number;
declare function tan(value: number): number;

$.time = {
    year,
    month,
    day,
    hour,
    minute,
    second,
    nanosecond,
    now
};

$.crypto = {
    md4,
    md5,
    sha1,
    sha224,
    sha512_224,
    sha256,
    sha512_256,
    sha384,
    sha512,
    keccak256,
    keccak512,
    sha3_224,
    sha3_256,
    sha3_384,
    sha3_512
};

$.hash = {
    adler32,
    fnv1_128,
    fnv1a_128,
    fnv1_32,
    fnv1a_32,
    fnv1_64,
    fnv1a_64
};

$.math = {
    abs,
    acos,
    ceil,
    floor,
    round,
    randomInt,
    randomInt32,
    randomInt64,
    randomFloat32,
    randomFloat64,
    gamma,
    NaN,
    max,
    min,
    pow,
    sqrt,
    pi,
    piAsString,
    e,
    eAsString,
    phi,
    phiAsString,
    ln2,
    ln2AsString,
    ln10,
    ln10AsString,
    sqrt2,
    sqrt2AsString,
    log2E,
    log10E,
    maxInt8,
    minInt8,
    maxInt16,
    minInt16,
    maxInt32,
    minInt32,
    sqrtPi,
    sqrtPiAsString,
    log,
    log10,
    inf,
    isNaN,
    cosh,
    j0,
    j1,
    erf,
    erfinv,
    erfc,
    erfcinv,
    exp,
    exp2,
    expm1,
    hypot,
    ilogb,
    y0,
    y1,
    tanh,
    trunc,
    tan  
};