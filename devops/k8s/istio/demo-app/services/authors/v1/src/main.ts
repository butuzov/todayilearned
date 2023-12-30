import { NestFactory } from '@nestjs/core';
import { createLightship } from 'lightship';

import { AppModule } from './app.module';
import { ValidationPipe } from '@nestjs/common';


const PORT = process.env.PORT || 3000

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const lightship = createLightship();

  lightship.registerShutdownHandler(() => app.close());
  app.setGlobalPrefix('/api/v1');

  await app.listen(PORT);

  lightship.signalReady();
  return app.getUrl();
}

bootstrap();
