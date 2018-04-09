create table tags (
    name text primary key
);

create table command_tags (
    command integer, 
    tag text,
    primary key (command, tag),
    foreign key(command) references command(id),
    foreign key(tag) references tags(name)
);