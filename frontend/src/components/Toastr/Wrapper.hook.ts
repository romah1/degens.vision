import { useAppSelector } from '../../store/hooks';
import { getAllMessages } from './redux';

export const WrapperHook = () => {
    const messages = useAppSelector(getAllMessages);

    return {
        messages,
    };
};
