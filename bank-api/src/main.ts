import { NestFactory } from '@nestjs/core';
import { Transport } from '@nestjs/microservices';
import { AppModule } from './app.module';
import { ModelNotFoundExceptionFilter } from './exception-filters/model-not-found.exception-filter';

async function bootstrap() {
  const app = await NestFactory.create(AppModule, {cors: true});
  app.setGlobalPrefix('api'); // um prefixo global para todas as rotas
  app.useGlobalFilters(new ModelNotFoundExceptionFilter()); // Para os nossos exception filters.
  app.connectMicroservice({ //Configurações de consumer para se comunicar com o Kafka.
    transport: Transport.KAFKA,
    options: {
      client: {
        brokers: [process.env.KAFKA_BROKER],
      },
      consumer: {
        groupId:
          !process.env.KAFKA_CONSUMER_GROUP_ID ||
          process.env.KAFKA_CONSUMER_GROUP_ID === ''
            ? 'my-consumer-' + Math.random()
            : process.env.KAFKA_CONSUMER_GROUP_ID,
      },
    },
  });
  app.startAllMicroservices();
  await app.listen(3000);
}
bootstrap();
