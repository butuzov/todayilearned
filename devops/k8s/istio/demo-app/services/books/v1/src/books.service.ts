import { Injectable } from '@nestjs/common';

import { Book } from './books.interface';
import { booksRepository } from './books.repository';

@Injectable()
export class AppService {
  private readonly books: Map<string, Book> = booksRepository;

  public getBooks(): Book[] {
    return Array.from(this.books.values());
  }

  public findById(id: string): Book {
    return this.books.get(id);
  }
}
