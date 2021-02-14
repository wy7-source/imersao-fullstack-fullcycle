import { Console, Command } from "nestjs-console";
import { getConnection } from "typeorm";
import * as chalk from 'chalk'; // Para colorir o console.log
// É a implementação de um comando para automatizar a população do banco.
@Console()
export class FixturesCommand{

    @Command({
        command: 'fixtures',
        description: 'Seed data in database'
    }) // Decorator para especificar que isso se trata de um comando.
    async command(){
        await this.runMigrations();
        const fixtures = (await import(`./fixtures/bank-${process.env.BANK_CODE}`)).default // Carregamos os arquivos de fixtures.
        for(const fixture of fixtures){
            //Inserimos no banco cada fixture
            await this.createInDatabase(fixture.model, fixture.fields);
        }

        console.log(chalk.green('Data generated'));
    }

    async runMigrations(){
        const conn = getConnection('default'); // Carrega as variáveis de ambiente para pegar a conexão.
        for(const migration of conn.migrations.reverse()){
            // Iremos reverter a ordem das migrações em que foram criadas para desfazer o banco.
            await conn.undoLastMigration();
        }
    }

    async createInDatabase(model: any, data: any){
        const repository = this.getRepository(model);
        const obj = repository.create(data);
        await repository.save(obj);
    }

    getRepository(model: any){
        // Para pegarmos o repositório presente na model.
        const conn = getConnection('default');
        return conn.getRepository(model);
    }
}