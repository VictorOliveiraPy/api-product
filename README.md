



# Estruturas de pastas

``/cmd``
Principais aplicações para este projeto.

O nome do diretório para cada aplicação deve corresponder ao nome do executável que você deseja ter (ex. /cmd/myapp).

Não coloque muitos códigos no diretório da aplicação. Se você acha que o código pode ser importado e usado em outros projetos, ele deve estar no diretório /pkg. Se o código não for reutilizável ou se você não quiser que outros o reutilizem, coloque esse código no diretório /internal. Você ficará surpreso com o que os outros farão, então seja explícito sobre suas intenções!

É comum ter uma pequena função main que importa e invoca o código dos diretórios /internal e /pkg e nada mais.

Veja o diretório /cmd para mais exemplos.

``/internal``
Aplicação privada e código de bibliotecas. Este é o código que você não quer que outras pessoas importem em suas aplicações ou bibliotecas. Observe que esse padrão de layout é imposto pelo próprio compilador Go. Veja o Go 1.4 release notes para mais detalhes. Observe que você não está limitado ao diretório internal de nível superior. Você pode ter mais de um diretório internal em qualquer nível da árvore do seu projeto.

Opcionalmente, você pode adicionar um pouco de estrutura extra aos seus pacotes internos para separar o seu código interno compartilhado e não compartilhado. Não é obrigatório (especialmente para projetos menores), mas é bom ter dicas visuais que mostram o uso pretendido do pacote. Seu atual código da aplicação pode ir para o diretório /internal/app (ex. /internal/app/myapp) e o código compartilhado por essas aplicações no diretório /internal/pkg (ex. /internal/pkg/myprivlib).

``/pkg``
Código de bibliotecas que podem ser usados por aplicativos externos (ex. /pkg/mypubliclib). Outros projetos irão importar essas bibliotecas esperando que funcionem, então pense duas vezes antes de colocar algo aqui :-) Observe que o diretório internal é a melhor maneira de garantir que seus pacotes privados não sejam importáveis porque é imposto pelo Go. O diretório /pkg contudo é uma boa maneira de comunicar explicitamente que o código naquele diretório é seguro para uso. I'll take pkg over internal A postagem no blog de Travis Jeffery fornece uma boa visão geral dos diretórios pkg e internal, e quando pode fazer sentido usá-los.

É também uma forma de agrupar o código Go em um só lugar quando o diretório raiz contém muitos componentes e diretórios não Go, tornando mais fácil executar várias ferramentas Go (conforme mencionado nestas palestras: Best Practices for Industrial Programming da GopherCon EU 2018, GopherCon 2018: Kat Zien - How Do You Structure Your Go Apps e GoLab 2018 - Massimiliano Pippi - Project layout patterns in Go).

Consulte o diretório /pkg se quiser ver quais repositórios Go populares usam esse padrão de layout de projeto. Este é um padrão de layout comum, mas não é universalmente aceito e alguns na comunidade Go não o recomendam.

Não há problema em não usá-lo se o projeto do seu aplicativo for muito pequeno e onde um nível extra de aninhamento não agrega muito valor (a menos que você realmente queira :-)). Pense nisso quando estiver ficando grande o suficiente e seu diretório raiz ficar muito ocupado (especialmente se você tiver muitos componentes de aplicativos não Go).

``/api``
Especificações OpenAPI/Swagger, arquivos de esquema JSON, arquivos de definição de protocolo.

``/test``
Aplicações de testes externos adicionais e dados de teste. Sinta-se à vontade para estruturar o diretório /test da maneira que quiser. Para projetos maiores, faz sentido ter um subdiretório de dados. Por exemplo, você pode ter /test/data ou /test/testdata se precisar que o Go ignore o que está naquele diretório. Observe que o Go também irá ignorar diretórios ou arquivos que começam com "." ou "_", para que você tenha mais flexibilidade em termos de como nomear seu diretório de dados de teste.

```Gerar swagger```
swag init -g cmd/server/main.go 

