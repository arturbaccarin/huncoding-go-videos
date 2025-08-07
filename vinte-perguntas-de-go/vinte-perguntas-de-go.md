# 20 Perguntas de Go para Entrevistas Sênior

Youtube: https://youtu.be/lw5W_QLhrLk

## 1.  Como funciona o scheduler de goroutines no runtime do Go?

Como ele cria as gorotinas dentro do projeto, quando você coloca a palavra Go antes do método.

O Go utiliza o modelo m:n, sendo o m o número de gorotinas e n o número de threads do SO para esse projeto.

O scheduler gerencia a utilização de recursos de gorotinas fazendo com que ele consiga suspender a execução de algumas gorotinas para otimizar o uso de threads dentro do SO.

Preempção preemptiva e preempção cooperativa são duas formas de controle de execução de processos em um sistema operacional. 

Na preempção preemptiva, o sistema operacional pode interromper um processo em execução a qualquer momento para dar lugar a outro, garantindo maior controle e divisão de tempo entre os processos. 

Já na preempção cooperativa, o processo em execução é responsável por liberar voluntariamente o processador, ou seja, ele precisa "cooperar" com os outros, o que pode causar problemas se ele travar ou não liberar o controle. A principal diferença entre elas é quem decide quando a troca ocorre: o sistema (preemptiva) ou o próprio processo (cooperativa).

## 2. Quais são os impactos do Garbage Collector em sistemas de baixa latência? Como mitigar esses impactos?

Embora o Garbage Collector (GC) do Go seja altamente otimizado, ele ainda pode causar pequenas pausas durante a coleta de lixo, o que pode impactar negativamente sistemas que exigem baixa latência. Para mitigar esses impactos, é possível ajustar o comportamento do GC por meio da variável de ambiente GOGC, que permite controlar a frequência da coleta de lixo — valores mais altos reduzem a frequência, mas aumentam o uso de memória. Além disso, é importante evitar alocações dinâmicas desnecessárias, especialmente dentro de loops críticos, e reduzir o uso de variáveis temporárias. Um bom gerenciamento de memória e o uso consciente dos recursos ajudam a manter a latência baixa e o desempenho estável.

GOGC é uma variável de ambiente que controla a frequência da coleta de lixo (GC - Garbage Collection) na linguagem Go. Ela determina quando o coletor de lixo deve ser executado, com base no crescimento da memória alocada.

O valor de GOGC representa uma percentagem de crescimento de memória desde a última coleta de lixo.

* Por padrão, GOGC = 100, o que significa que o GC será ativado quando a quantidade de memória usada duplicar (crescer 100%) desde a última coleta.
* Se você definir GOGC = 200, o GC será acionado apenas quando a memória crescer 200% (triplicar).
* Se GOGC = 50, o GC será acionado mais frequentemente (quando crescer apenas 50%).

O GC do Go é do tipo "concurrent and parallel mark-and-sweep", ou seja:

1. Concurrent (concorrente):
A maior parte do trabalho do GC é feita enquanto a aplicação ainda está rodando — sem precisar parar tudo.

2. Parallel (paralelo):
Ele usa múltiplas threads para coletar lixo ao mesmo tempo, o que melhora o desempenho em sistemas com vários núcleos.

3. Mark-and-sweep (marcar e limpar):
* Mark (marcar): O GC identifica os objetos que ainda estão sendo usados.
* Sweep (limpar): Remove da memória os objetos que não estão mais sendo usados.

O GC do Go minimiza as pausas, mas ainda precisa parar brevemente a aplicação em certos momentos, principalmente no início da fase de "mark" (quando ele identifica os objetos ativos). Essas pausas são chamadas de "STW" — Stop The World.

* Evite criar muitos objetos durante a execução, principalmente dentro de loops ou partes do código muito usadas.

* Reutilize estruturas de dados sempre que possível, em vez de criar novas toda vez.

* Reduza o uso de variáveis temporárias, que são criadas e descartadas rapidamente, pois geram trabalho extra para o GC.

* Planeje bem o uso de memória, reservando o espaço necessário antecipadamente para evitar realocações.

* Menos alocações = menos trabalho para o Garbage Collector, o que significa menos pausas e melhor desempenho.

## 3. Como você faria profiling e benchmarking para identificar gargalos em um programa concorrente em Go?

Entra num ponto de performance que muitas vezes você tem que fazer no código.

Basicamente pergunta como fazer um profiling, fazer um benchmark no seu código e verificar gargalos da sua aplicação, pelo menos inicialmente com a ferramenta.

No momento que você cria um arquivo _test, você consegue criar funções teste e funções benchmark. E assim você já consegue testar performance por ele.

Pacote profiling: vai verificar quanto de memória ele está utilizando em alta latência, em alta carga dentro do seu projeto. Ou reetorna o uso de CPU.

## 4. O que são race conditions e como o Go te ajuda a identificá-los?

Race conditions basicamente é quando você tem múltiplas go routines acessando um dado compartilhado.

Comando go run -race e ele vai identificar quando tem race condition ali na execução.

Consegue identificar que tem algum ponto, ele fala em qual linha possíveis race condition e assim por diante.

---
Pergunta cinco. Como implementar um pool de grutines eficiente e quais são as boas práticas
06:13
para isso? Então pool aqui de grotines é para você não criar grotine infinitamente, né? Então não quero que o meu código saia criando 1 milhão de grutine, porque grutin são sim leves, mas elas não são infinitas também. Então a gente tem que ter um limite ali e uma pool geralmente é utilizada para você fazer essa limitação, tá? Então você consegue fazer um uma sincronia de mensagens ali utilizando channels com buffer. Então é um dos jeitos de você fazer isso. Você pode utilizar o workers, então em vez de você criar 1
06:39
milhão de grutines, você cria grutines que eu vi em channels, eles viram workers. É como se você tivesse criando um lista, uma fila ali dentro do seu código. E aí você consegue já gerenciar isso melhor. Você consegue criar menos grutines, né? Então você consegue fazer uma alimentação. Com os workers você consegue fazer um limite. Então quando você cria workers, você escala de uma forma com um limite. Então fala lá, escale 20 workers, 30 workers e não infinitos workers. Então fica melhor. Beleza? Pergunta seis. Como funciona o
07:07
escape analysis no compilador Go e qual seu impacto na performance? Tá? Então, como você consegue fazer análise de escape dentro do Gol e como isso afeta ou como isso pode te ajudar na questão de performance? Então, um compilador basicamente ele consegue decidir se uma variável vai ser alocada na stack, então na stack ali de memória para poder guardar essa variável ou se ela vai ser jogada pra HIP. Então, geralmente essas variáveis que que vão pra HIP, elas geralmente dão um overhead ali no garbage collector, tá? Então ela acaba
07:33
estourando o garbage collector porque toda hora tem hip, tem coisa para ele limpar, assim dizendo. Então, basicamente o GC ele tenta fazer um gerenciamento do que ele deve enviar ou não para HIP, porque quanto mais coisas no HIP, mais trabalho ele tem depois para poder fazer a limpeza do lixo, assim, vamos dizer, né? Então, basicamente, só para entender esse funcionamento de quando uma variável ou não é utilizada, como o garbage collector trabalha, aqui é uma pergunta mais rasa, mas é uma pergunta que pode
07:58
ser importante de ser feito, tá? Pergunta sete: Quais são as armadilhas comuns ao utilizar canais com buffer e como evitá-las? Essa daqui é muito boa, tá? Então, a gente tem os canais com buffer e sem buffer. Vou deixar, tem vídeo aqui no canal também, vou deixar aqui no card, beleza? Então, os canais com buffer quando você abre uma memória para poder colocar vários valores naquele channel. Quando não tem buffer, você basicamente coloca um valor de um para um. Então, eu só escrevo se alguém tiver lendo e só leio se alguém tiver
08:23
escrevendo. Beleza? Então, se você começar a ler num channel que não tem buffer, está vazio, você vai ficar preso lendo para sempre. E se você tá tentando escrever num buffer que não tem ninguém ouvindo, então você vai est escrevendo e esperando escrever para sempre, porque o buffer não tem listener, tá? Essa é a diferença. Agora, quando ele tem buffer, você consegue escrever e sair. Quando não tem buffer, você consegue ler e esperar alguém escrever, tá? Então, essas são as diferenças de buffer e com
08:48
um channel sem buffer. Porém, se você tiver um buffer, um canal com buffer lotado, então você tem lá um buffer de 100, você escreveu 100 caracteres, 100 itens lá dentro e você tentar escrever mais um item, você vai criar um deadlock, porque não vai ter ninguém ouvindo nesse channel, tem lá 100 posições sem ninguém ouvir e você tá tentando escrever mais um ainda. Então você não criou ninguém para ler e a pessoa que tá escrevendo, o channel, a grutine que tá escrevendo, não consegue escrever porque lotou. Então você vai
09:15
criar um deadlock. Então outro problema é você criar um channel com 1 milhão de posições no buffer e não utilizar. E aí você tá criando uma memória, uma utilização de memória gigante, uma reserva, uma reserva de memória gigante e você nunca vai utilizar isso, tá? Então você tá subutilizando o channel com buffer enorme que você acabou criando. Então você tem que gerenciar melhor esse buffer e alguns problemas clássicos são os próprios deadlocks no momento que você tá utilizando o buffer com channel, tá? ou channel com buffer
09:41
nesse caso. Pergunta oito. Explica o funcionamento do tipo interface vazia. O interface abre e fechaves em go e seus impactos de performance. Então, como vocês já sabem, o interface abre fechaves ou N que foi substituído, na verdade foi criado um alias no GO nas novas versões. Basicamente ele é um formato de você escrever qualquer tipo lá dentro, tá? Então, basicamente ele recebe qualquer tipo. Por isso que ele foi criado um eles de N para ficar mais parecido com as outras linguagens e mais familiar, né? Só que tem um problema.
10:08
Quando você utiliza esse tipo interface ou um N, você tá criando basicamente um overhead de valores, porque você tá colocando um valor gigante numa variável que pode ser representada como tipo valor, tá? Então o uso excessivo dele pode gerar aquele box unboxing que basicamente você tá colocando um valor X dentro da interface, mas o go tem que toda hora descobrir qual é essa tipagem ou você tem que ficar fazendo casting e migrações dentro do seu código para poder pegar o tipo certo. Você tem que
10:32
ficar utilizando, por exemplo, Generics para você saber generics para todo lado. Você não sabe qual tipagem, você tem que fazer casting, fazer validação se aquele tipo é o realmente que você espera que tá dentro daquele daquela variável. Então você vai criando vários problemas por utilizar a interface. Claro que muitas vezes não tem como não utilizar. Por exemplo, quando você recebe um Jason que você não quer mapear todos os campos, você pode criar um map de string N ou string interface, mas a gente pelo
10:59
menos entender o o problema que ele pode causar, o problema de performance já é muito bom, tá? Pergunta nove. Como você estruturaria um sistema de microsserviços em Go para garantir escabilidade e resiliência? Então aqui como você cria um sistema de microsserviço para garantir que você escale ou ele fique resiliente, que ele remova recursos, que ele realmente não fique com recursos executando depois de terminar e assim por diante. Então aqui a gente pode entrar no contexto do context, tá? Que é muito poderoso aqui
11:25
em go. Então a gente pode utilizar context para colocar deadline, para utilizar use cancel para poder matar grotines que estão executando. Assim a gente melhora a questão de performance e uso de recursos. A gente pode utilizar trace Ability para colocar logs, open telemetry, por exemplo, para colocar B3CD com a questão do Zipkin, por exemplo. A gente pode colocar secretaker, rate limits nos routers para poder implementar rate limit, secret breaker dentro do seu projeto para para ele parar de ficar tentando várias vezes
11:54
quando ele perceber que é um problema. A gente pode utilizar mensagens assíncronas e filas ao invés de colocar métodos diretos para poder fazer requisições diretas. Aí você causa um problema em massa porque você não tem utilização e uma resiliência no seu projeto. Pessoa escrevo várias mensagens, elas todas me tornam 500. Depois eu não consigo reprocessar. Então com filas você consegue trabalhar melhor isso. Aqui entra um pouco de arquitetura, mas é um bom pensamento ali para trazer para uma entrevista, beleza?
12:19
Ou que pode te trazer caso você tiver participando como candidato. Pergunta 10: qual é o modelo de memória em GO? como garantir visibilidade e evitar riscondition, tá? Então entra um pouco com aquela outra pergunta que basicamente aqui vai entrar um pouquinho. Basicamente a pergunta é como evita riscondition. Aqui a forma de evitar são várias, mas a em geral se a pessoa se você conseguir responder Mtex geralmente já vai ser muito bom, tá? Claro que ela vai estender muito mais do que isso, porque Mutex pode causar
12:46
questão de lock novamente, questão de dead lock, porque nunca vai soltar aquele MTEX ou você começa a adquirir tanto Miltex, tanto Miltex, que você trava sua aplicação inteira. E ela nunca vai soltar porque você travou todo mundo, todas as variáveis com Mtex, né? Tem um vídeo aqui de lock free também que eu gravei no canal que também fala um pouquinho sobre formatos para você não utilizar Mtex em todo canto na sua aplicação. Então tem outros formatos ali para você utilizar, mas o Mtex em geral vai responder bem essa pergunta, tá?
13:12
Então a questão de operações atômicas, se você tiver utilizando contadores, por exemplo, utilizar Atomic, o pacote Atomic do Go, tem várias respostas a gente pode utilizar aqui nessa pergunta. Beleza? Pergunta 11. como lidar com Penix e recuperar a execução do projeto. Então aqui a resposta simples, é o recover, tá? Geralmente você coloca o recover com defer dentro do seu método e quando ele termina de executar, ele percebe que deu um penic, ele cai nesse recover e você consegue recuperar o código sem ele derrubar. Algumas
13:41
bibliotecas, como por exemplo Jingonic, ele já executa isso com recover, tá? Então você já consegue ter um router com recover. Então, se quebrar lá no seu repository com Penic, ele vai quebrar no repository, no service, no router, vai quebrar e vai o dingonic, ele vai segurar essa execução para você e não vai derrubar o seu projeto, tá? Então aqui a gente consegue utilizar e responder basicamente com defer com o recover aqui. Beleza? Pergunta 12. Qual a diferença entre métodos com receiver por valor e por ponteiro? Então, quais
14:11
as consequências no comportamento do programa? Então você tem um receiver ali num método e não função. Método. Quando você coloca ali function, abre e fecha parênteses e coloca o nome da struct. Ali ser um ponteiro e ali não ser um ponteiro. Qual era qual é a diferença dele? Tá? A diferença aqui, essa pergunta é até simples, mas a diferença basicamente é você ter um objeto que quando você alterar ele não vai alterar o objeto original, vai ser uma cópia. Então basicamente você criar métodos com ponteiro, evita cópia, evita de você
14:37
ficar enviando um monte de objeto e cada um ser uma cópia diferente um do outro. Isso causa inconsistência no seu código. Às vezes vai ser difícil você pegar até em logs ou em produção. Então é importante sempre que você criar um método de receiver, você colocar ele como ponteiro, principalmente, na verdade obrigatoriamente, se tiver no mesmo pacote, porque o go obriga, tá? Então aqui essa pergunta se responde dessa forma, basicamente. É uma pergunta simples, mas que pode ser uma pegadinha. Beleza? Pergunta 13. Como interfaces são
15:03
implementadas implicitamente em GO e como isso afeta a arquitetura do código? Então, basicamente, se eu tenho uma interface que ele tem vários métodos, como eu faço a implementação de forma implícita dessa interface em Go. Então, basicamente, se um tipo definir todos os métodos e ele basicamente tá implementando isso, o go implementa essa interface automaticamente com a struct. Inclusive, é algo muito difícil de alguém que veio de outras linguagens entender. Porque lá no Java você coloca basicamente o implements, o Super, você
15:30
tem várias formas de você poder fazer isso e implementar uma interface, mas no GO ele não é feito de forma explícita. você não diz que aquele que aquela struct está implementando aquela interface, ele é feito de forma implícita e ele é feito basicamente quando você coloca uma structando todos os métodos. Então, basicamente ele faz isso de forma implícita. Você implementa todos os métodos e assim ele sabe que aquela interface tá sendo implementada por alguém. Isso no Golend, por exemplo, na ideia dá para ver bem fácil com
15:56
aquele I verde que aparece na esquerda ou no VS Code você consegue ir pra implementação e ele vai cair na struct que implementa todos os métodos, tá? Então, basicamente é assim que a gente responde também uma pergunta um pouco simples aqui. Pergunta 14. Como funcionam os channels em go e quais são as principais formas de sincronizar routines com elas? Tá? Essa pergunta é muito boa, porque a gente consegue entender basicamente o channel como fifo, né? Então o channel no fim ele é uma fila, então ele basicamente pode ser
16:25
utilizado como um SQS da WS, por exemplo, dentre as grotines, porque a ordem que foi enviada as mensagens é a ordem que elas vão ser lidas também, tá? Então aqui no caso first in first out, beleza? Então enviar e receber dados são com buffer não bloqueantes, tá? Agora, se você coloca um buffer cheio ou você coloca um channel sem buffer, ele é bloqueante, sim. Então o gerenciamento aqui de growes com channels, ele pode ser feito utilizando buffer, só que com cuidado, porque entra naquele primeiro
16:52
ponto de buffer que a gente falou aqui, né? Então, basicamente, a gente consegue utilizar chanciar e enviar mensagens de uma G routine para outra, sinalizando coisa, sinalizando erro, sinalizando o sucesso e assim por diante. Então, a gente consegue gerenciar ou então gerenciando o tempo. Por exemplo, o context utiliza channel para fazer o deadline, para fazer cancelamento. Então, ele também utiliza basicamente channel para poder fazer envio de mensagens e de sinais entre as grotines. Então, é um formato ali de você utilizar
17:18
o gerenciamento do channel para gerenciar as grutines que estão executando no projeto. Beleza? Pergunta 15. Para que serve o pacote Context e como ele pode ser usado para gerenciar deadlines, cancelamentos e propagação de dados? Entra um pouco com uma outra pergunta. Então, basicamente, o context ele é imutável. A partir do momento que você cria um contexto, ele não muda. Não muda mais, a não ser que você crie um contexto a partir desse original. Então, no momento que você cria um contexto com cancelamento e você executa o cancel,
17:44
todas as pessoas que receberam contexto por parâmetro vão receber essa notificação de cancelamento. Por isso que em go é importante você receber o context em todos os métodos como primeiro parâmetro, tá? De forma o o padrão ali do goxto seja sempre o primeiro parâmetro, tá? Então, como ele é imutável, todo mundo que recebeu o contexto, mesmo que tenha criado um contexto em cima desse que ele recebeu, ele vai receber essa notificação de cancel e é assim que ele basicamente gerencia ali o cancelamento e o deadline
18:12
dentro do seu projeto, enviando uma mensagem e garantindo que todo mundo vai receber ela, porque todos todas as pessoas ali, todas as partes do seu código tem o mesmo valor do Miltex imutável, mesmo que tenha criado o outro em cima dele. Beleza? Pergunta 16. Como funciona o mecanismo de Differ? em que situações ele é crucial para evitar leakso, tá? Então o defer aqui, como se você não conhece, basicamente é uma palavrinha que você coloca no go que ele sempre vai executar no fim do seu código, tá? Então você coloca defer,
18:41
quando tudo terminar depois do return ele vai executar essa função que você colocou no defer, tá? Vale lembrar, se você quiser deixar essa pergunta mais difícil, você vale lembrar o seguinte: qual é a ordem que o defer cria? Então se eu colocar três de fur, um embaixo do outro, quem vai executar primeiro? o primeiro que eu coloquei de fur ou o último, tá? E e aqui a resposta basicamente é que o defer é um lifo, tá? Então o último que você dor é o primeiro que vai ser executado quando o método terminar. Então ele vai basicamente
19:09
ajudar você com link de recurso ou evitar líquid de recurso, porque no momento que eu cria conexão com banco de dados, eu posso logo em seguida colocar um defer para desligar ela. Ou então eu crio um teste contêiner no meu teste, eu posso logo em seguida colocar um defer para remover esse contêiner. Então, evita que eu esqueça de colocar esse código ou coloque numa parte do código que nunca vai chegar. Às vezes eu coloquei um return antes de executar o código que remove esse contêiner, que tira a conexão. Então eu posso
19:35
basicamente retornar ele e colocar ele para evitar que isso aconteça. Beleza? Pergunta 17. Como implementar a injeção de dependências em GO para garantir ter estabilidade de baixo acoplamento? Então, como a gente pode implementar uma injeção de dependência em GO eficiente? Basicamente isso. Em gestão do Gol a gente sabe que é um pouco humorosa, que a gente precisa ficar criando métodos construtores que retornam uma struct. Essa structar um por parâmetro do outro, por parâmetro do outro. Então aqui a gente pode
20:05
utilizar parâmetros e construtores se a gente quiser fazer isso na mão. E basicamente a gente cria construtores dentro de cada parte do código, service, repositor, us case, controller e assim por diante. E aí a gente consegue utilizar isso de forma manual, porém a gente consegue utilizar bibliotecas como Google Wire, por exemplo. Então o Google Wire ele vai basicamente utilizar reflection go para saber quais variáveis já foram inicializadas e quais construtores precisam daquelas variáveis que já foram inicializadas para serem
20:31
enviadas por parâmetro. Beleza? Então a gente consegue responder aqui dessa forma. Algo simples aqui, mas só pra gente saber como funciona, tá? Pergunta 18. Como o pacote Sync pode ser utilizado para gerenciar concorrência e quais são os principais permitíveis que ele oferece, tá? Então aqui a gente tem um pacote Sc e a gente tá basicamente perguntando quais são as partes principais que a gente consegue utilizar dele. Então a gente tem basicamente o 11, o condition e o H group aqui que são os principais fora o Mtex que a gente já
20:56
conversou. Então o Mutex muito importante para evitar aqui da gente utilizar a mesma variável, ficar inconsistente dentro do nosso código, tá? A gente tem o W group, que é para basicamente notificar várias grutines de que algo tem que ser terminado. Então você cria 20 grutines. Ao invés de você criar um tempo randômico para esperar as 20, você recebe 20 sinais de dano e aí você consegue terminar o seu código sabendo que as 20 grotines terminaram de executar. Aí você tem o 11, que ele basicamente vai fazer uma inicialização
21:24
única. Então, ao invés de eu garantir que várias partes do meu código estão criando uma inicialização de uma variável, eu consigo colocar um 11. E independente do que aconteça, vai inicializar apenas uma vez dentro do seu código, tá? E você tem um conde para enviar notificações em massa aqui para várias grouines, beleza? Como se a gente tivesse utilizando o channel aqui também. Então, a gente tem esses principais pacotes ali, essas principais funções dentro do pacote syc. Beleza? Então, como o Gold trata operações de
21:49
tempo, isso aqui é muito importante também, e quais são os principais padrões para usá-lo corretamente? Então o pacote time aqui ele oferece o ticker e o after. Então aqui o after basicamente como o nome já diz, ele fala depois de tanto tempo envie uma notificação ou faça algo então time ponto after ele basicamente vira um channel que te retorna uma mensagem quando esse determinado tempo passar. Geralmente você coloca ele dentro de um select e você recebe uma notificação ali e você faz o que tem que fazer quando
22:16
esse tempo terminar de passar. Beleza? Já o ticker ele é uma função do time, do after que nesse caso, só que ele vai enviar a mesma notificação a cada tempo. Então o after a cada 20 segundos ele vai te enviar uma mensagem e vai parar. O ticker você vai colocar 20 segundos. Então a cada 20 segundos ele vai fazer alguma coisa para você. Então ah, a cada 20 segundos me envie uma mensagem. A cada 20 segundos coloque um logger. Aí você vai colocar os dois num select. E aí quando algo acontecer no after, ele
22:41
vai executar apenas uma vez e pronto. No ticker, a cada x tempos, ele vai executar aquilo que você pedir para ele. Beleza? Então ele vai ficar ficar fazendo tics de tempo para você, vamos assim dizer. Beleza? E última pergunta. Quais as melhores práticas para escrever testes concorrentes em go e evitar falsos positivos ou negativos? Então, no GO você tem o testem ponto t e dentro dessa variável t você consegue ter o t. Perlol, beleza? Então ele vai executar basicamente seus códigos de forma concorrente aqui, de forma paralela aqui
23:10
tem executados ao mesmo tempo, tá? E aí, como você garante que você está evitando falas positivo, riscondition, você consegue executar com menos race os seus testes. Aí você garante que não tem riscondition entre os seus testes também, beleza? Isso no go é possível. E você consegue isolar os dados dessas variáveis, porque como tem vários testes executando ao mesmo tempo, se você colocar uma variável compartilhada, vão estar todos os testes alterando ao mesmo tempo. Então você consegue isolar, coloque variáveis de de contexto local
23:37
dentro dos testes para que você evite de que você tenha ali uma uma variável compartilhada que tá ficando inconsistente entre as execuções concorrentes dos testes. Beleza? Então essas seriam aqui 20 perguntas. Fica um vídeo extenso, mas fica fica aí um um conteúdo diferente para vocês, para vocês conhecerem e talvez estudar coisas diferentes aí no dia a dia de vocês. Eu vou trazer vários desses temas aqui no canal, principalmente sobre concorrência, é garbage collector e compilador aqui do Gol, beleza? Mas isso
24:06
daqui a pouco a gente vai ver um pouquinho mais aqui no canal. Então vejo vocês na próxima e até lá. Falou. [Música]