import {
  BeforeInsert,
  Column,
  CreateDateColumn,
  Entity,
  JoinColumn,
  ManyToOne,
  PrimaryGeneratedColumn,
} from 'typeorm';
import { BankAccount } from './bank-account.model';
import { v4 as uuidv4 } from 'uuid';

export enum PixKeyKind {
  cpf = 'cpf',
  email = 'email',
}

@Entity({ name: 'pix_keys' })
export class PixKey {
  @PrimaryGeneratedColumn('uuid')
  id: string;

  // Podemos colocar os valores possíveis para essa variável, por uma enum.
  @Column()
  kind: PixKeyKind;

  @Column()
  key: string;

  // Relacionamento clássico de FK.
  @ManyToOne(() => BankAccount)
  @JoinColumn({ name: 'bank_account_id' })
  bankAccount: BankAccount;

  // Para não precisarmos buscar a conta, e depois atribui-la na criação da PixKey. 
  @Column()
  bank_account_id: string;

  @CreateDateColumn({ type: 'timestamp' })
  created_at: Date;

  @BeforeInsert() generateId() {
    if (this.id) {
      return;
    }
    this.id = uuidv4();
  }
}