export interface UserInitResponse {
    nonce: string;
}

export interface UserVerifyRequest {
    address: string;
    signature: string;
    ref_code?: string;
}

export interface UserVerifyResponse {
    token: string;
    expire: string;
    code: number;
}
