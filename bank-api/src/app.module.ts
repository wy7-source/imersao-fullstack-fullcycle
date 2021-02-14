import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { MyFirstController } from './controllers/my-first/my-first.controller';
import { BankAccount } from './models/bank-account.model';
import { BankAccountController } from './controllers/bank-account/bank-account.controller';
import { ConsoleModule } from 'nestjs-console';
import { FixturesCommand } from './fixtures/fixtures.command';
import { PixKeyController } from './controllers/pix-key/pix-key.controller';
import { PixKey } from './models/pix-key.model';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { TransactionController } from './controllers/transaction/transaction.controller';
import { Transaction } from './models/transaction.model';
import { TransactionSubscriber } from './subscribers/transaction-subscriber/transaction-subscriber.service';

@Module({
  imports: [
    // forRoot, importa todos os artefatos do módulo para o AppModule raiz em específico.
    ConfigModule.forRoot(), //Módulo para leitura de variáveis de ambiente.
    ConsoleModule, // Nosso módulo para habilitar comandos de console.
    TypeOrmModule.forRoot({ 
      // As variáveis de ambiente que foram lidas pelo ConfigModule.
      type: process.env.TYPEORM_CONNECTION as any,
      host: process.env.TYPEORM_HOST,
      port: parseInt(process.env.TYPEORM_PORT),
      username: process.env.TYPEORM_USERNAME,
      password: process.env.TYPEORM_PASSWORD,
      database: process.env.TYPEORM_DATABASE,
      entities: [BankAccount, PixKey, Transaction] // Entidades disponíveis.
    }),
    // forFeatura, importa para uso na aplicação.
    TypeOrmModule.forFeature([BankAccount, PixKey, Transaction]), // Entidades Habilitadas para uso na aplicação.
    //Configurações para se comunicar com o CodePix por GRPC.
    ClientsModule.register([
      {
        name: 'CODEPIX_PACKAGE',
        transport: Transport.GRPC,
        options: {
          url: process.env.GRPC_URL,
          package: 'github.com.codeedu.codepix', // O mesmo nome de pacote que temos no protofile do CodePix.
          protoPath: [join(__dirname, 'protofiles/pixkey.proto')] // Para importar os ProtoFiles.
        }
      }
    ]),
    ClientsModule.register([
      //Configurações de producer para se comunicar com o Kafka.
      {
        name: 'TRANSACTION_SERVICE',
        transport: Transport.KAFKA,
        options: {
          client: {
            clientId: process.env.KAFKA_CLIENT_ID,
            brokers: [process.env.KAFKA_BROKER]
          },
          consumer: {
            groupId: !process.env.KAFKA_CONSUMER_GROUP_ID ||
              process.env.KAFKA_CONSUMER_GROUP_ID === ''
                ? 'my-consumer-' + Math.random()
                : process.env.KAFKA_CONSUMER_GROUP_ID,
          }
        }
      }
    ])
  ],
  controllers: [AppController, MyFirstController, BankAccountController, PixKeyController, TransactionController],
  providers: [AppService, FixturesCommand, TransactionSubscriber],
})
export class AppModule {}
