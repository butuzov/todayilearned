import {
  Controller,
  Get,
  Param,
  UseInterceptors,
  ClassSerializerInterceptor,
  Query,
  ValidationPipe,
} from '@nestjs/common';
import { ApiProperty } from '@nestjs/swagger';

import { AppService } from './authors.service';
import { Author } from './authors.interface';
import { Transform } from 'class-transformer';

export class IFiltersDTO {
  @Transform(id => id.split(','))
  ids?: string[];
}
@Controller('authors')
@UseInterceptors(ClassSerializerInterceptor)
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getAuthors(
    @Query(new ValidationPipe({ transform: true })) filters: IFiltersDTO,
  ): Author[] {
    console.log('filters', filters);
    return this.appService.getAuthors();
  }

  @Get(':id')
  getById(@Param('id') id: string): Author {
    return this.appService.findById(id);
  }
}
