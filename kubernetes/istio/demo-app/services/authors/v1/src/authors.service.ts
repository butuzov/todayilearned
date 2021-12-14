import { Injectable } from '@nestjs/common';

import { Author } from './authors.interface';
import { authorsRepository } from './authors.repository';

@Injectable()
export class AppService {
  private readonly authors: Map<string, Author> = authorsRepository;

  public getAuthors(): Author[] {
    return Array.from(this.authors.values());
  }

  public findById(id: string): Author {
    return this.authors.get(id);
  }
}
