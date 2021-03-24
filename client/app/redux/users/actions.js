import {
    CREATE_USER_REQUEST,
    GET_USER_LIST_REQUEST
} from './constants';

export const userListRequest = () => ({
    type: GET_USER_LIST_REQUEST,
});

export const createUserRequest = (email, first_name, second_name, password, repeat_password, role = 'moderator') => ({
    type: CREATE_USER_REQUEST,
    email: email,
    first_name: first_name,
    second_name: second_name,
    password: password,
    repeat_password: repeat_password,
    role: role
});

