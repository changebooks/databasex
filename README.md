# databasex
database + log
==

<pre>
profile, err := log.NewProfile(map[string]string{
    log.ProfileDirectory: "./log",
    log.ProfileChannel:   "database",
})

if err != nil {
    fmt.Println(err)
    return
}

stream, err := log.NewStream(profile)
if err != nil {
    fmt.Println(err)
    return
}

logger, err := log.NewLogger(stream, "test", 1)
if err != nil {
    fmt.Println(err)
    return
}

idRegister := &log.IdRegister{}
idRegister.SetTraceId("trace-id-10001")
idRegister.SetBizId("biz-id-20002")
</pre>

<pre>
write, err := database.NewProfile(map[string]string{
    database.ProfileId:       "test",
    database.ProfileDriver:   "mysql",
    database.ProfileHost:     "127.0.0.1",
    database.ProfileDatabase: "test",
    database.ProfileUsername: "root",
    database.ProfilePassword: "123456",
    database.ProfileWrite:    "true",
})

if err != nil {
    fmt.Println(err)
    return
}

builder := database.DriversBuilder{}

err = builder.AddProfile(write)
if err != nil {
    fmt.Println(err)
    return
}

drivers, err := builder.Build()
if err != nil {
    fmt.Println(err)
    return
}

driver, err := drivers.GetWriter()
if err != nil {
    fmt.Println(err)
    return
}

db, err := databasex.New(logger)
if err != nil {
    fmt.Println(err)
    return
}
</pre>

<pre>
query := `DROP TABLE IF EXISTS user;`
_, err11 := db.Exec(idRegister, driver, query)
if err11 != nil {
    fmt.Println(err11)
    return
}

query = `CREATE TABLE user (
    id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL DEFAULT '' COMMENT '姓名',
    phone varchar(32) NOT NULL DEFAULT '' COMMENT '手机号',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`

_, err12 := db.Exec(idRegister, driver, query)
if err12 != nil {
    fmt.Println(err12)
    return
}
</pre>

<pre>
r21, err21 := db.Exec(idRegister, driver, "INSERT INTO user (name, phone) VALUES (?, ?);", "张三", "13000000001")
if err21 != nil {
    fmt.Println(err21)
    return
}

fmt.Println(r21.LastInsertId())

r22, err22 := db.Exec(idRegister, driver, "INSERT INTO user (name, phone) VALUES (?, ?);", "李四", "13000000002")
if err22 != nil {
    fmt.Println(err22)
    return
}

fmt.Println(r22.LastInsertId())
</pre>

<pre>
r31, err31 := db.Exec(idRegister, driver, "UPDATE user SET phone = ? WHERE id = ?", "18000000001", 1)
if err31 != nil {
    fmt.Println(err31)
    return
}

fmt.Println(r31)
</pre>

<pre>
r41, err41 := db.First(idRegister, driver, "SELECT * FROM user WHERE id = ?", 1)
if err41 != nil {
    fmt.Println(err41)
    return
}

fmt.Println(r41)

r42, err42 := db.Find(idRegister, driver, "SELECT * FROM user")
if err42 != nil {
    fmt.Println(err42)
    return
}

fmt.Println(r42)

r43, err43 := db.AggregateInt(idRegister, driver, "SELECT COUNT(1) AS 'aggregate' FROM user")
if err43 != nil {
    fmt.Println(err43)
    return
}

fmt.Println(r43)
</pre>

<pre>
// 进程正常关闭前
err51 := drivers.Close()
fmt.Println(err51)
</pre>
