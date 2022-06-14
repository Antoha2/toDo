DROP TABLE todolist;

create table todolist (
	task_id serial not null,
	user_id	int not null,
	Text varchar(255) not null,
	IsDone bool not null default false	,
	primary key(task_id),
	foreign key(user_id) references userlist (user_id)
)