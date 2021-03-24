import {
    GET_USER_LIST_ERROR,
    GET_USER_LIST_REQUEST,
    GET_USER_LIST_SUCCESS,
} from './constants';

const initialState = {
    isLoading: false,
    successful: false,
    error: '',
    users: [],
};

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case GET_USER_LIST_REQUEST:
            return {
                ...state,
                isLoading: true,
            };
        case GET_USER_LIST_SUCCESS:
            return {
                error: '',
                isLoading: false,
                successful: true,
                users: action.payload,
            };
        case GET_USER_LIST_ERROR:
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
