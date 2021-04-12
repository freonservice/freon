import {
    LIST_IDENTIFIERS_REQUEST,
    CREATE_IDENTIFIER_REQUEST,
    UPDATE_IDENTIFIER_REQUEST,
    DELETE_IDENTIFIER_REQUEST,
} from './constants';

export const listIdentifiersRequest = () => ({
    type: LIST_IDENTIFIERS_REQUEST,
});

export const createIdentifierRequest = (name, description = '', example_text = '', categoryId = 0, platforms = [], namedList = []) => ({
    type: CREATE_IDENTIFIER_REQUEST,
    name: name,
    description: description,
    example_text: example_text,
    platforms: platforms,
    namedList: namedList,
    categoryId: categoryId,
});

export const updateIdentifierRequest = (id, name, description = '', example_text = '', categoryId = 0, platforms = [], namedList = []) => ({
    type: UPDATE_IDENTIFIER_REQUEST,
    id: id,
    name: name,
    description: description,
    example_text: example_text,
    platforms: platforms,
    namedList: namedList,
    categoryId: categoryId
});

export const deleteIdentifierRequest = (id) => ({
    type: DELETE_IDENTIFIER_REQUEST,
    id: id,
});
