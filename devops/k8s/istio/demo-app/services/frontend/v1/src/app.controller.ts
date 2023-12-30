import { Controller, Get } from '@nestjs/common';
import { Observable } from 'rxjs';

import { AppService } from './app.service';

@Controller('dashboard')
export class AppController {
  constructor(private readonly appService: AppService) {}

  @Get()
  getDashboard(): Observable<Object> {
    return this.appService.getDashboard();
  }
}
