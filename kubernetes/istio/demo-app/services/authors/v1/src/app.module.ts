import { Module } from '@nestjs/common';
import { AppController } from './authors.controller';
import { AppService } from './authors.service';

@Module({
  imports: [],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
