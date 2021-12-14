import { Author } from './authors.interface';

export const authorsRepository = new Map<string, Author>([
  [
    '1',
    {
      id: '1',
      firstName: 'Loreth Anne',
      lastName: 'White',
      born: 'South Africa',
    },
  ],
  [
    '2',
    {
      id: '2',
      firstName: 'Lisa',
      lastName: 'Regan',
      born: 'USA',
    },
  ],
  [
    '3',
    {
      id: '3',
      firstName: 'Ty',
      lastName: 'Patterson',
      born: 'USA',
    },
  ],
  [
    '4',
    {
      id: '4',
      firstName: 'Sue',
      lastName: 'Burke',
      born: 'USA',
    },
  ],
  [
    '5',
    {
      id: '5',
      firstName: 'Aliya',
      lastName: 'Whiteley',
      born: 'UK',
    },
  ],
  [
    '6',
    {
      id: '6',
      firstName: 'Yoon Ha',
      lastName: 'Lee',
      born: 'USA',
    },
  ],
]);
