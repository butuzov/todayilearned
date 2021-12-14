// local env
// export const authorsBaseUrl = 'http://localhost:3002/api/v1';
// export const booksBaseUrl = 'http://localhost:3001/api/v1';

export const authorsBaseUrl = process.env.URL_BASE_AUTHORS || 'http://authors:3000/api/v1';
export const booksBaseUrl = process.env.URL_BASE_BOOKS || 'http://books:3000/api/v1';
