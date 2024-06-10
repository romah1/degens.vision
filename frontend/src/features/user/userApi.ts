import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

import { baseApiUrl, defaultFetchParams } from '../../config';
import { UserInitResponse, UserVerifyRequest, UserVerifyResponse } from '../../interfaces/user';

export const userApi = createApi({
    reducerPath: 'userApi',
    baseQuery: fetchBaseQuery({ ...defaultFetchParams, baseUrl: `${baseApiUrl}/auth/` }),
    endpoints: builder => ({
        userInit: builder.mutation<UserInitResponse, string>({
            query: publicKey => ({
                url: 'init',
                method: 'POST',
                body: {
                    address: publicKey,
                },
            }),
        }),
        userVerify: builder.mutation<UserVerifyResponse, UserVerifyRequest>({
            query: body => ({
                url: 'verify',
                method: 'POST',
                body,
            }),
        }),
    }),
});

export const {
    useUserInitMutation,
    useUserVerifyMutation,
} = userApi;

export const { endpoints, reducerPath, reducer, middleware } = userApi;
