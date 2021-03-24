export const apiServerURL = 'http://localhost:4000';

export function handleApiErrors(response) {
    if (!response.ok) throw new Error(response.statusText);
    return response;
}