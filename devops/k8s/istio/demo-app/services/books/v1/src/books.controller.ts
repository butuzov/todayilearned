import {
  Controller,
  Get,
  Param,
  UseInterceptors,
  ClassSerializerInterceptor,
} from '@nestjs/common';

import { AppService } from './books.service';
import { Book } from './books.interface';

@Controller('books')
@UseInterceptors(ClassSerializerInterceptor)
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getBooks(): Book[] {
    return this.appService.getBooks();
  }

  @Get(':id')
  getById(@Param('id') id: string): Book {
    return this.appService.findById(id);
  }
}
