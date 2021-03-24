import {
    STAT_ERROR,
    STAT_REQUEST,
    STAT_SUCCESS
} from './constants';

const initialState = {
    isLoading: false,
    successful: false,
    error: '',
    stat: {},
};

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case STAT_REQUEST:
            return {
                ...state,
                isLoading: true,
            };
        case STAT_SUCCESS:
            return {
                error: '',
                isLoading: false,
                successful: true,
                stat: action.payload,
            };
        case STAT_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        default:
            return state;
    }
}
