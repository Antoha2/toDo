DROP TABLE todolist ;

create table todolist (
	task_id serial primary key not null,
	Text varchar(255) not null,
	IsDone bool not null default false	
)