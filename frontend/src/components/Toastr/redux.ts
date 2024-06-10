import { createEntityAdapter, createSlice, PayloadAction } from '@reduxjs/toolkit';
import { Message } from './interfaces';
import { RootState } from '../../store';

export const toastrAdapter = createEntityAdapter<Message>({
    selectId: message => message.id,
});

const initialState = toastrAdapter.getInitialState();

export const toastrSlice = createSlice({
    name: 'toastrMessages',
    initialState,
    reducers: {
        addToastrMessage: (state, payload: PayloadAction<Omit<Message, 'id'>>) => {
            const id = Number(new Date());
            state.ids.push(id);
            state.entities[id] = {
                id,
                detail: payload.payload.detail,
                level: payload.payload.level,
                title: payload.payload.title,
            };
        },
        removeToastrMessage: toastrAdapter.removeOne,
    },
});

export const {
    addToastrMessage,
    removeToastrMessage,
} = toastrSlice.actions;

export const {
    selectAll: getAllMessages,
} = toastrAdapter.getSelectors<RootState>(state => state.toastrMessages);
