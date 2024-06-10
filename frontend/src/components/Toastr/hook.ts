import React from 'react';

import { FetchBaseQueryError } from '@reduxjs/toolkit/dist/query';
import { SerializedError } from '@reduxjs/toolkit';
import { useAppDispatch } from 'src/store/hooks';
import { addToastrMessage } from 'src/components';

interface CustomError {
    data: {
        message: string;
    }
}

export function useErrorToaster(error: FetchBaseQueryError | SerializedError | undefined, title: string = 'Backend Error', errorText: string = 'Something went wrong, please try again later') {

    const dispatch = useAppDispatch();
    React.useEffect(
        () => {
            if (!error) {
                return;
            }
            const detail = errorText;
            dispatch(addToastrMessage({
                level: 'error',
                title,
                detail,
            }));
        },
        [dispatch, error, title],
    );
}
