import {
    CREATE_LOCALIZATION_ERROR, CREATE_LOCALIZATION_SUCCESS,
    LIST_LOCALIZATIONS_ERROR,
    LIST_LOCALIZATIONS_REQUEST,
    LIST_LOCALIZATIONS_SUCCESS
} from './constants';

const initialState = {
    error: '',
    isLoading: false,
    successful: false,
    listLocalizations: []
};

export default function reducer(state = initialState, action) {
    switch (action.type) {
        case LIST_LOCALIZATIONS_REQUEST:
            return {
                ...state,
                isLoading: true,
            };
        case LIST_LOCALIZATIONS_SUCCESS:
            return {
                error: '',
                isLoading: false,
                successful: true,
                listLocalizations: action.payload,
            };
        case LIST_LOCALIZATIONS_ERROR:
            return {
                ...state,
                error: action.error.toString(),
                isLoading: false,
                successful: false,
            };
        case CREATE_LOCALIZATION_SUCCESS:
            return {
                ...state,
                error: '',
                isLoading: false,
                successful: true,
            };
        case CREATE_LOCALIZATION_ERROR:
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
