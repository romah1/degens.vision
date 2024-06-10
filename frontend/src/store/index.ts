import { configureStore } from '@reduxjs/toolkit';
import { collectionsApi } from '../features/collections/collectionsApi';
import { userApi } from '../features/user/userApi';
import { toastrSlice } from '../components';

export const store = configureStore({
    reducer: {
        [toastrSlice.name]: toastrSlice.reducer,
        [collectionsApi.reducerPath]: collectionsApi.reducer,
        [userApi.reducerPath]: userApi.reducer,
    },
    middleware: getDefaultMiddleware =>
        getDefaultMiddleware().concat([
            collectionsApi.middleware,
            userApi.middleware,
        ]),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
