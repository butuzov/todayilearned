import { Module } from '@nestjs/common';
import { AppController } from './books.controller';
import { AppService } from './books.service';

@Module({
  imports: [],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
