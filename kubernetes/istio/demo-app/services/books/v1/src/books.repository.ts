import { Book } from './books.interface';

export const booksRepository = new Map<string, Book>([
  [
    '1',
    {
      id: '1',
      title: 'Semiosis: A Novel',
      pages: 333,
      authorId: '4',
      genres: [
        'Novel',
        'Science Fiction',
        'Space opera',
        'Hard science fiction',
      ],
    },
  ],
  [
    '2',
    {
      id: '2',
      title: 'The Loosening Skin',
      pages: 132,
      authorId: '5',
      genres: ['Science Fiction'],
    },
  ],
  [
    '3',
    {
      id: '3',
      title: 'Ninefox Gambit',
      pages: 384,
      authorId: '6',
      genres: ['Novel', 'Science Fiction'],
    },
  ],
  [
    '4',
    {
      id: '4',
      title: 'Raven Stratagem',
      pages: 400,
      authorId: '6',
      genres: ['Science Fiction', 'Fantasy Fiction'],
    },
  ],
  [
    '5',
    {
      id: '5',
      title: 'Revenant Gun',
      pages: 466,
      authorId: '6',
      genres: ['Science Fiction', 'Adventure fiction'],
    },
  ],
]);
