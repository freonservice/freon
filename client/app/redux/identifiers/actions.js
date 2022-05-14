import {
    LIST_IDENTIFIERS_REQUEST,
    CREATE_IDENTIFIER_REQUEST,
    UPDATE_IDENTIFIER_REQUEST,
    DELETE_IDENTIFIER_REQUEST,
} from './constants';

export const listIdentifiersRequest = () => ({
    type: LIST_IDENTIFIERS_REQUEST,
});

export const createIdentifierRequest = (name, description = '', text_singular = '', text_plural = '', categoryId = 0, platforms = [], namedList = []) => ({
    type: CREATE_IDENTIFIER_REQUEST,
    name: name,
    description: description,
    text_singular: text_singular,
    text_plural: text_plural,
    platforms: platforms,
    namedList: namedList,
    categoryId: categoryId,
});

export const updateIdentifierRequest = (id, name, description = '', text_singular = '', text_plural = '', categoryId = 0, platforms = [], namedList = []) => ({
    type: UPDATE_IDENTIFIER_REQUEST,
    id: id,
    name: name,
    description: description,
    text_singular: text_singular,
    text_plural: text_plural,
    platforms: platforms,
    namedList: namedList,
    categoryId: categoryId
});

export const deleteIdentifierRequest = (id) => ({
    type: DELETE_IDENTIFIER_REQUEST,
    id: id,
});
