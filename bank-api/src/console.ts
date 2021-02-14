import { BootstrapConsole } from "nestjs-console";
import { AppModule } from "./app.module";
// Arquivo para iniciar o AppModule, independente de compilação, será inteiramente em typescript.
const bootstrap = new BootstrapConsole({
    module: AppModule,
    useDecorators: true
})

bootstrap.init().then(async app => {
    // Para termos a instância da aplicação.
    try{
        await app.init();
        await bootstrap.boot();
        process.exit(0);
    }catch(e){
        console.error(e);
        process.exit(1);
    }
})