CREATE TABLE roles
(
	role_id    serial primary key,
	rolesname varchar(255) not null unique
)

CREATE TABLE userlist_to_roles
(
	user_id int not null,
	role_id int not null,
	primary key(user_id, role_id),
	foreign key(user_id) references userlist (user_id),
	foreign key(role_id) references roles(role_id)
)
