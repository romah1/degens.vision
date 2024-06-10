import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

import { baseApiUrl, defaultFetchParams } from '../../config';
import { TopMitsResponse, CollectionInfoResponse, LiveMintsItem } from '../../interfaces/collections';

export const collectionsApi = createApi({
    reducerPath: 'collectionsApi',
    baseQuery: fetchBaseQuery({
        ...defaultFetchParams,
        baseUrl: `${baseApiUrl}/api/v1/`,
        prepareHeaders: (headers) => {
            const sessionToken = localStorage.getItem('token');
            if (sessionToken) {
                headers.set('Authorization', `Bearer ${sessionToken}`)
            }
            return headers;
        }
    }),
    endpoints: builder => ({
        getCollectionsTop: builder.mutation<TopMitsResponse, number>({
            query: minutesGap => `collections/top?minutes_gap=${minutesGap}`,
        }),
        getCollectionInfo: builder.mutation<CollectionInfoResponse, string>({
            query: collection_key => `collections/full?collection_key=${collection_key}`,
        }),
        getInitLiveFeed: builder.query<LiveMintsItem[], void>({
            query: () => 'mints/initial_feed',
        }),
    }),
});

export const {
    useGetCollectionsTopMutation,
    useGetCollectionInfoMutation,
    useGetInitLiveFeedQuery,
} = collectionsApi;

export const { endpoints, reducerPath, reducer, middleware } = collectionsApi;
