package utils

const (
	INSERT_MAHASISWA             = `INSERT INTO m_mahasiswa values (?,?,?,?)`
	SELECT_MAHASISWA             = `SELECT * FROM m_mahasiswa`
	UPDATE_MAHASISWA             = `UPDATE m_mahasiswa SET nama_mahasiswa=?,jenis_kelamin=?,alamat=? where nim = ?`
	DELETE_MAHASISWA             = `DELETE from m_mahasiswa where nim=?`
	SELECT_MAHASISWA_WITH_MATKUL = `SELECT 
    table_mahasiswa_matkul.nim,
    table_mahasiswa_matkul.nama_mahasiswa,
    table_mahasiswa_matkul.jenis_kelamin,
    table_mahasiswa_matkul.alamat,
    table_mahasiswa_matkul.kode_matkul,
    table_mahasiswa_matkul.nilai,
    m_mata_kuliah.nama_matkul,
    m_mata_kuliah.kode_dosen
	FROM
    (SELECT 
        m.nim,
            m.nama_mahasiswa,
            m.jenis_kelamin,
            m.alamat,
            mm.kode_matkul,
            mm.nilai
    FROM
        m_mahasiswa m
    INNER JOIN m_mahasiswa_matkul mm ON m.nim = mm.nim) table_mahasiswa_matkul
        JOIN m_mata_kuliah ON table_mahasiswa_matkul.kode_matkul = m_mata_kuliah.kode_matkul`
	INSERT_DOSEN  = `INSERT INTO m_dosen values (?,?)`
	SELECT_DOSEN  = `SELECT * FROM m_dosen`
	UPDATE_DOSEN  = `UPDATE m_dosen SET nama_dosen=? where kode_dosen = ?`
	DELETE_DOSEN  = `DELETE from m_dosen where kode_dosen=?`
	INSERT_MATKUL = `INSERT INTO m_mata_kuliah VALUES(?,?,?)`
	SELECT_MATKUL = `SELECT 
    m_mata_kuliah.kode_matkul,
    m_mata_kuliah.nama_matkul,
    m_mata_kuliah.kode_dosen,
    m_dosen.nama_dosen
    FROM
    m_mata_kuliah
        INNER JOIN
    m_dosen ON m_mata_kuliah.kode_dosen = m_dosen.kode_dosen`
	UPDATE_MATKUL = `UPDATE m_mata_kuliah
    SET kode_dosen = ?
    where kode_matkul = ?`
	DELETE_MATKUL           = `DELETE FROM m_mata_kuliah where kode_matkul=?`
	INSERT_MAHASISWA_MATKUL = `INSERT INTO m_mahasiswa_matkul(nim,kode_matkul,nilai) VALUES(?,?,?)`
	SELECT_MAHASISWA_MATKUL = `SELECT 
    m_mahasiswa_matkul.nim,
    m_mahasiswa_matkul.kode_matkul,
    m_mata_kuliah.nama_matkul,
    m_mata_kuliah.kode_dosen,
    m_mahasiswa_matkul.nilai
FROM
    m_mahasiswa_matkul
        INNER JOIN
    m_mata_kuliah ON m_mahasiswa_matkul.kode_matkul = m_mata_kuliah.kode_matkul where nim=?`
	UPDATE_MAHASISWA_MATKUL = `UPDATE m_mahasiswa_matkul
    set nilai = ?,nilaiNumber = ?
    where kode_matkul = ? and nim=?`

	SELECT_MATKUL_BY_KODE = `SELECT 
    m_mahasiswa_matkul.nim,
    m_mahasiswa_matkul.kode_matkul,
    m_mata_kuliah.nama_matkul,
    m_mata_kuliah.kode_dosen,
    m_mahasiswa_matkul.nilai
FROM
    m_mahasiswa_matkul
        INNER JOIN
    m_mata_kuliah ON m_mahasiswa_matkul.kode_matkul = m_mata_kuliah.kode_matkul where m_mahasiswa_matkul.kode_matkul=?`

	SELECT_REPORT_MATKUL = `SELECT 
    m_mahasiswa_matkul.kode_matkul,
    m_mata_kuliah.nama_matkul,
    m_mata_kuliah.kode_dosen,
    avg(m_mahasiswa_matkul.nilaiNumber) as rata_rata
    FROM
    m_mahasiswa_matkul
        INNER JOIN
    m_mata_kuliah ON m_mahasiswa_matkul.kode_matkul = m_mata_kuliah.kode_matkul
    group by m_mahasiswa_matkul.kode_matkul `
)
