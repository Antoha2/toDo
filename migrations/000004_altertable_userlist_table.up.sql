DROP TABLE todolist;
DROP TABLE userlist;

create table userlist (
	user_id serial PRIMARY KEY not null,
	firstname varchar(255) not null,
	lastname varchar(255) not null,
	username varchar(255) not null,
	password varchar(255) not null	
)

create table todolist (
	task_id serial not null,
	user_id	int not null,
	Text varchar(255) not null,
	IsDone bool not null default false	,
	primary key(task_id),
	foreign key(user_id) references userlist (user_id)
)