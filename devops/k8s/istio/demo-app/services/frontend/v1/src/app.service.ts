import { Injectable, HttpService } from '@nestjs/common';
import { Observable, forkJoin } from 'rxjs';
import { map } from 'rxjs/operators';

import { authorsBaseUrl, booksBaseUrl } from './app.constants';

@Injectable()
export class AppService {
  constructor(private httpService: HttpService) {}

  getDashboard(): Observable<Object> {
    return forkJoin({
      authors: this.getAuthors(),
      books: this.getBooks()
    });
  }

  getAuthors(): Observable<any> {
    return this.httpService
      .get(`${authorsBaseUrl}/authors/`)
      .pipe(map(v => v.data));
  }

  getBooks(): Observable<any> {
    return this.httpService.get(`${booksBaseUrl}/books/`).pipe(map(v => v.data));
  }
}
