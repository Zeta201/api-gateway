package com.zeta.university.reviewservicev2.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.zeta.university.reviewservicev2.entity.Review;

public interface ReviewRepository extends JpaRepository<Review, Long>{

}
